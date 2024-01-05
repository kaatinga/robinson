package robinson

// FunctionNotPassedError is an error type for when a required function is not passed.
type FunctionNotPassedError struct{}

func (e FunctionNotPassedError) Error() string {
	return "function not passed"
}

func (e FunctionNotPassedError) Is(err error) bool {
	_, ok := err.(FunctionNotPassedError)
	return ok
}

// NewFunctionNotPassedError creates a new instance of FunctionNotPassedError.
func NewFunctionNotPassedError() error {
	return FunctionNotPassedError{}
}
