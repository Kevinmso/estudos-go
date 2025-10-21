package dberrors

import "fmt"

type NotFound struct {
	Entity string
	ID     string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("entity %s with id %s not found", e.Entity, e.ID)
}
