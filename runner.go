package runner

import (
	"context"
	"sync/atomic"
	"time"
)

type TaskInterface interface {
	Do() func() error
}

type Task struct {
	Work func() error
}

func (c *Task) Do() func() error {
	return c.Work
}

type Runner struct {
	queue    *queue
	maxNum   int64
	curNum   int64
	interval time.Duration
	ctx      context.Context

	// execute task error
	OnError func(err error)

	// syscall.SIGINT/SIGTERM, maybe you should backup data
	OnClose func(data []TaskInterface)

	// wait for OnClose done
	Wait chan bool
}

// num: max concurrent number
// interval: interval of checking new task
func New(ctx context.Context, num int64, interval time.Duration) *Runner {
	o := &Runner{
		queue:    newQueue(),
		maxNum:   num,
		curNum:   0,
		interval: interval,
		ctx:      ctx,
		Wait:     make(chan bool),
	}
	go o.run()

	return o
}

func (c *Runner) run() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if atomic.LoadInt64(&c.curNum) >= c.maxNum {
				continue
			}

			if t, ok := c.queue.front(); ok {
				atomic.AddInt64(&c.curNum, 1)
				go func() {
					err := t.Do()()
					atomic.AddInt64(&c.curNum, -1)
					if err != nil && c.OnError != nil {
						c.OnError(err)
					}
				}()
			}
		case <-c.ctx.Done():
			if c.OnClose != nil {
				c.OnClose(c.queue.clear())
			}
			c.Wait <- true
			return
		}
	}
}

func (c *Runner) Add(t TaskInterface) {
	c.queue.push(t)
}
