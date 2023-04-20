package validator

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) AddError(key, value string) {
	if _, ok := v.Errors[key]; !ok {
		v.Errors[key] = value
	}
}

func (v *Validator) Check(ok bool, key, value string) {
	if !ok {
		v.AddError(key, value)
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
