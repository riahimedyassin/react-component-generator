package generator

// Validating the args, type etc.
type Validator interface {
	Validate() error
}
