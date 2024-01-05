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

// Get returns current value from cache.
func (c *Crusoe[ValueType]) Get() ValueType {
	c.RLock()
	output := c.value
	c.RUnlock()
	return output
}

// Set sets value to cache.
func (c *Crusoe[ValueType]) Set(value ValueType) {
	c.Lock()
	c.value = value
	c.Unlock()
}

// Call calls function with current value and using arbitrary processing in the function may set a new value to cache.
// It is possible to use this function to update cache value using an external source or whatever, the operation is atomic.
func (c *Crusoe[ValueType]) Call(f func(v ValueType) ValueType) {
	if f == nil {
		return
	}
	c.Lock()
	defer c.Unlock()
	c.value = f(c.value)
}

// CallWithError calls function with current value and allows error handling.
// The operation is atomic. If the function returns an error, the cache value is not updated.
func (c *Crusoe[ValueType]) CallWithError(f func(v ValueType) (ValueType, error)) error {
	if f == nil {
		return NewFunctionNotPassedError()
	}
	c.Lock()
	defer c.Unlock()
	newValue, err := f(c.value)
	if err != nil {
		return err
	}
	c.value = newValue
	return nil
}
