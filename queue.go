package runner

import "sync"

type Queue struct {
	mutex *sync.Mutex
	data  []interface{}
}

func newQueue() *Queue {
	return &Queue{
		mutex: &sync.Mutex{},
		data:  make([]interface{}, 0),
	}
}

func (c *Queue) Push(eles ...interface{}) {
	c.mutex.Lock()
	c.data = append(c.data, eles...)
	c.mutex.Unlock()
}

func (c *Queue) pop() interface{} {
	c.mutex.Lock()

	length := len(c.data)
	var result interface{}
	switch length {
	case 0:
		result = nil
	case 1:
		result = c.data[0]
		c.data = make([]interface{}, 0)
	default:
		result = c.data[0]
		c.data = c.data[1:]
	}

	c.mutex.Unlock()
	return result
}

func (c *Queue) Len() int {
	c.mutex.Lock()
	length := len(c.data)
	c.mutex.Unlock()
	return length
}

func (c *Queue) clear() []interface{} {
	c.mutex.Lock()
	data := c.data
	c.data = make([]interface{}, 0)
	c.mutex.Unlock()
	return data
}
