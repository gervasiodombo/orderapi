package shared

import "github.com/google/uuid"

type UUIDGeneratorImpl struct{}

func New() *UUIDGeneratorImpl {
	return &UUIDGeneratorImpl{}
}

func (g *UUIDGeneratorImpl) Generate() string {
	return uuid.New().String()
}
