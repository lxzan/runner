package runner

import "sync"

type queue struct {
	mutex *sync.Mutex
	data  []TaskInterface
}

func newQueue() *queue {
	return &queue{
		mutex: &sync.Mutex{},
		data:  make([]TaskInterface, 0),
	}
}

func (c *queue) push(t TaskInterface) {
	c.mutex.Lock()
	c.data = append(c.data, t)
	c.mutex.Unlock()
}

func (c *queue) front() (t TaskInterface, ok bool) {
	c.mutex.Lock()

	length := len(c.data)
	if length == 1 {
		t = c.data[0]
		ok = true
		c.data = make([]TaskInterface, 0)
	} else if length > 1 {
		t = c.data[0]
		ok = true
		c.data = c.data[1:]
	} else {
		ok = false
	}

	c.mutex.Unlock()
	return
}

func (c *queue) len() int {
	c.mutex.Lock()
	length := len(c.data)
	c.mutex.Unlock()
	return length
}

func (c *queue) clear() []TaskInterface {
	c.mutex.Lock()
	data := c.data
	c.data = make([]TaskInterface, 0)
	c.mutex.Unlock()
	return data
}
