package gateway

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type EncoderGateway struct{ cost int }

func New() *EncoderGateway {
	return &EncoderGateway{cost: bcrypt.DefaultCost}
}

func (e *EncoderGateway) Encode(rawPassword string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(rawPassword), e.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashed), nil
}
