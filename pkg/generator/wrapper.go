package generator

// Wrapper is an interface for generation structs.
type Wrapper interface {
	GetDefiners() []FileSpecDefiner
}
