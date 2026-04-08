package gateway_test

import (
	"testing"

	"github.com/oderapi/infra/gateway"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorIfBcryptFails(t *testing.T) {
	//Arrange
	encoder := &gateway.EncoderGatewayImpl{Cost: 100} //invalid cost
	//Act
	_, err := encoder.Encode("secret")
	//Assert
	assert.NotNil(t, err, "expected no error, got none")
	assert.Contains(t, err.Error(), "failed to hash password")
}
