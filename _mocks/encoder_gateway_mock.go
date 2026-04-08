package _mocks

type EncoderGatewayMock struct {
	EncodeResult string
	EncodeErr    error
	EncodeCalled bool
	EncodeParam  string
}

func (m *EncoderGatewayMock) Encode(rawPassword string) (string, error) {
	m.EncodeCalled = true
	m.EncodeParam = rawPassword
	return m.EncodeResult, m.EncodeErr
}
