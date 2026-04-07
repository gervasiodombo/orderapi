package gateway

import (
	"errors"
)

type EncoderGateway interface {
	Encode(rawPassword string) (string, error)
}

var ErrEncoder = errors.New("an error occurred while encoding the password")
