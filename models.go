package robinson

// Crusoe is single-value cache model.
type Crusoe[ValueType any] struct {
	value    ValueType
	setValue chan ValueType
	getValue chan chan ValueType
}

// NewCrusoe creates new single-value cache item.
func NewCrusoe[ValueType any]() *Crusoe[ValueType] {
	crusoe := &Crusoe[ValueType]{
		setValue: make(chan ValueType),
		getValue: make(chan chan ValueType),
	}

	// Launch the cache.
	go crusoe.runCache()

	return crusoe
}

func (c *Crusoe[ValueType]) runCache() {
	for {
		select {
		case cacheValue := <-c.setValue: // set value
			c.value = cacheValue
		case output := <-c.getValue: // read value
			output <- c.value
			close(output)
		}
	}
}

func (c *Crusoe[ValueType]) Get() ValueType {
	outputChannel := make(chan ValueType)
	c.getValue <- outputChannel
	return <-outputChannel
}

func (c *Crusoe[ValueType]) Set(value ValueType) {
	c.setValue <- value
}
