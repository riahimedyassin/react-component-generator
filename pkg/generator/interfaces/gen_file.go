package generator_interfaces

// Marker interface to mark files. This is usefull in generic function like clean up.
//
// See [generator_interfaces.Wrapper].
type GenFiles interface {
	MarkGenFile()
}
