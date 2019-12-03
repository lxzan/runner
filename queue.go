package runner

import "sync"

type Task struct {
	Do func() error
}

type queue struct {
	mutex *sync.Mutex
	data  []*Task
}

func newQueue() *queue {
	return &queue{
		mutex: &sync.Mutex{},
		data:  make([]*Task, 0),
	}
}

func (c *queue) push(t *Task) {
	c.mutex.Lock()
	c.data = append(c.data, t)
	c.mutex.Unlock()
}

func (c *queue) front() (t *Task, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	length := len(c.data)
	if length == 1 {
		t = c.data[0]
		ok = true
		c.data = make([]*Task, 0)
	} else if length > 1 {
		t = c.data[0]
		ok = true
		c.data = c.data[1:]
	} else {
		ok = false
	}
	return
}

func (c *queue) len() int {
	c.mutex.Lock()
	length := len(c.data)
	c.mutex.Unlock()
	return length
}
