package runner

import (
	"context"
	"sync/atomic"
	"time"
)

type Runner struct {
	queue    *queue
	maxNum   int64
	curNum   int64
	interval time.Duration
	ctx      context.Context

	// syscall.SIGINT/SIGTERM, maybe you should backup data
	OnClose func(data []interface{})

	OnTask func(opt interface{})
}

// num: max concurrent number per second, <=1000
// interval: interval of checking new task
func New(ctx context.Context, num int64) *Runner {
	var interval = 1000 / num
	o := &Runner{
		queue:    newQueue(),
		maxNum:   num,
		curNum:   0,
		interval: time.Duration(interval) * time.Millisecond,
		ctx:      ctx,
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

			if opt, ok := c.queue.front(); ok {
				atomic.AddInt64(&c.curNum, 1)
				go func() {
					if c.OnTask != nil {
						c.OnTask(opt)
					}
					atomic.AddInt64(&c.curNum, -1)
				}()
			}
		case <-c.ctx.Done():
			if c.OnClose != nil {
				c.OnClose(c.queue.clear())
			}
		}
	}
}

// add a task
func (c *Runner) Add(opt interface{}) {
	c.queue.push(opt)
}
