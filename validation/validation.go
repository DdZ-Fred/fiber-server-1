package validation

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}
