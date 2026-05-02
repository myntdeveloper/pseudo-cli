package store

import "fmt"

type ErrNotFoundByName struct {
	Name string
}
type ErrNotFoundByTag struct {
	Tag string
}

func (e *ErrNotFoundByName) Error() string {
	return fmt.Sprintf("pseudo %q not found", e.Name)
}

func (e *ErrNotFoundByTag) Error() string {
	return fmt.Sprintf("tag %q not found", e.Tag)
}
