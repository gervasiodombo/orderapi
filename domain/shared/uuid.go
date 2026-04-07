package shared

import (
	"errors"
)

type IDGenerator interface {
	Generate() string
}

type ID struct {
	value string
}

var ErrEmptyID = errors.New("could not generate ID because value is empty")

func NewID(value string) (ID, error) {
	if value == "" {
		return ID{}, ErrEmptyID
	}
	return ID{value: value}, nil
}

func (id ID) Value() string {
	return id.value
}

func (id ID) Equals(other ID) bool {
	return id.value == other.value
}

func (id ID) String() string {
	return id.value
}
