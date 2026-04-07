package mocks

type IDGeneratorMock struct {
	Value string
}

func (m *IDGeneratorMock) Generate() string {
	return m.Value
}
