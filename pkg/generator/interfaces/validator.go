package generator_interfaces

// Validating the args, type etc.
type Validator interface {
	Validate() error
}
