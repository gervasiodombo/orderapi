package shared

type Validator struct {
	Err *DomainError
}

func (v *Validator) Check(val string, fn func(string) (string, *DomainError)) string {
	if v.Err != nil {
		return val
	}

	result, err := fn(val)
	if err != nil {
		v.Err = err
	}
	return result
}
