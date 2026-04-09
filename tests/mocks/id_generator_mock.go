package mocks

type IDGeneratorMock struct {
	GenerateResult string
	GenerateCalled bool
}

func (m *IDGeneratorMock) Generate() string {
	m.GenerateCalled = true
	return m.GenerateResult
}
