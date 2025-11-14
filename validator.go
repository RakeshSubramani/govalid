package govalid

// Validator is implemented by all validators (string, number, struct)
type Validator interface {
	Error() error
}
