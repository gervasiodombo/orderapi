package shared

import "github.com/google/uuid"

type UUIDGeneratorImpl struct{}

func NewIDGenerator() *UUIDGeneratorImpl {
	return &UUIDGeneratorImpl{}
}

func (g *UUIDGeneratorImpl) Generate() string {
	return uuid.New().String()
}
