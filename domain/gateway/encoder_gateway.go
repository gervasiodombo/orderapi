package gateway

type EncoderGateway interface {
	Encode(rawPassword string) (string, error)
}
