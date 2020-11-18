package runner

import "sync"

type queue struct {
	mutex *sync.Mutex
	data  []interface{}
}

func newQueue() *queue {
	return &queue{
		mutex: &sync.Mutex{},
		data:  make([]interface{}, 0),
	}
}

func (c *queue) push(t interface{}) {
	c.mutex.Lock()
	c.data = append(c.data, t)
	c.mutex.Unlock()
}

func (c *queue) front() (t interface{}, ok bool) {
	c.mutex.Lock()

	length := len(c.data)
	switch length {
	case 0:
		ok = false
	case 1:
		t = c.data[0]
		ok = true
		c.data = make([]interface{}, 0)
	default:
		t = c.data[0]
		ok = true
		c.data = c.data[1:]
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

func (c *queue) clear() []interface{} {
	c.mutex.Lock()
	data := c.data
	c.data = make([]interface{}, 0)
	c.mutex.Unlock()
	return data
}
