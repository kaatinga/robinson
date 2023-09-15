package robinson

import "sync"

// Crusoe is single-value cache model.
type Crusoe[ValueType any] struct {
	value ValueType

	sync.RWMutex
}

// NewCrusoe creates new single-value cache item.
func NewCrusoe[ValueType any]() *Crusoe[ValueType] {
	return &Crusoe[ValueType]{}
}

func (c *Crusoe[ValueType]) Get() ValueType {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

func (c *Crusoe[ValueType]) Set(value ValueType) {
	c.Lock()
	defer c.Unlock()
	c.value = value
}

func (c *Crusoe[ValueType]) Call(f func(v ValueType) ValueType) {
	c.Lock()
	defer c.Unlock()
	c.value = f(c.value)
}
