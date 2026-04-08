package gateway

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type EncoderGatewayImpl struct{ Cost int }

func NewEncoderGateway() *EncoderGatewayImpl {
	return &EncoderGatewayImpl{Cost: bcrypt.DefaultCost}
}

func (e *EncoderGatewayImpl) Encode(rawValue string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(rawValue), e.Cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashed), nil
}
