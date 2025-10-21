package dberrors

type conflictError struct{}

func (e *conflictError) Error() string {
	return "attempted to create a record with an existing key"
}
