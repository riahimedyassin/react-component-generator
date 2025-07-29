package generator

type Transformer[T any] interface {
	GetTransformed() (*T, error)
}
