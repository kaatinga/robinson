package robinson

import "sync"

// Crusoe is single-value cache model.
type Crusoe[ValueType any] struct {
	value ValueType

	me sync.RWMutex
}

// NewCrusoePointer creates new single-value cache item and return a pointer to the item.
func NewCrusoePointer[ValueType any]() *Crusoe[ValueType] {
	return &Crusoe[ValueType]{}
}

// NewCrusoe creates new single-value cache item.
func NewCrusoe[ValueType any]() Crusoe[ValueType] {
	return Crusoe[ValueType]{}
}

// Get returns current value from cache.
func (c *Crusoe[ValueType]) Get() ValueType {
	c.me.RLock()
	output := c.value
	c.me.RUnlock()
	return output
}

// Set sets value to cache.
func (c *Crusoe[ValueType]) Set(value ValueType) {
	c.me.Lock()
	c.value = value
	c.me.Unlock()
}

// Call calls function with current value and using arbitrary processing in the function may set a new value to cache.
// It is possible to use this function to update cache value using an external source or whatever, the operation is atomic.
func (c *Crusoe[ValueType]) Call(f func(v ValueType) ValueType) {
	if f == nil {
		return
	}
	c.me.Lock()
	defer c.me.Unlock()
	c.value = f(c.value)
}

// CallWithError calls function with current value and allows error handling.
// The operation is atomic. If the function returns an error, the cache value is not updated.
func (c *Crusoe[ValueType]) CallWithError(f func(v ValueType) (ValueType, error)) error {
	if f == nil {
		return NewFunctionNotPassedError()
	}
	c.me.Lock()
	defer c.me.Unlock()
	newValue, err := f(c.value)
	if err != nil {
		return err
	}
	c.value = newValue
	return nil
}

// Check calls function with current value and returns the result of the function without changing the cache value.
func (c *Crusoe[ValueType]) Check(f func(v ValueType) bool) bool {
	if f == nil {
		return false
	}
	c.me.RLock()
	defer c.me.RUnlock()
	return f(c.value)
}
