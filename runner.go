package runner

import (
	"sync/atomic"
	"time"
)

type Runner struct {
	q        *queue
	maxNum   int64
	curNum   int64
	interval time.Duration
	OnError  func(err error)
}

// num: max concurrent number
// interval: interval of checking new task
func New(num int64, interval time.Duration) *Runner {
	o := &Runner{
		q:        newQueue(),
		maxNum:   num,
		curNum:   0,
		interval: interval,
	}
	go o.run()

	return o
}

func (c *Runner) run() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		num := atomic.LoadInt64(&c.curNum)
		if num >= c.maxNum {
			continue
		}

		if t, ok := c.q.front(); ok {
			atomic.AddInt64(&c.curNum, 1)
			go func() {
				err := t.Do()
				atomic.AddInt64(&c.curNum, -1)
				if err != nil && c.OnError != nil {
					c.OnError(err)
				}
			}()
		}
	}
}

func (c *Runner) Add(t *Task) {
	c.q.push(t)
}
