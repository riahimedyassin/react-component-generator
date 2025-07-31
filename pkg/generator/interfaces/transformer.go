package generator_interfaces

type Transformer[T any] interface {
	GetTransformed() (*T, error)
}
