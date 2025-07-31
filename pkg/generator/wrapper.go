package generator

// Wrapper is an interface for generation structs.
type Wrapper interface {
	// GetDefiners will let you get the file specs noted returned by struct implementing the FileSpecDefiner interface.
	// The GetDefiners function will get those definers based on the configuration and the flags passed in.
	GetDefiners() []FileSpecDefiner
	// GetEditors will let your get the files that will be edited. The core will be responsible for managing and updating those files.
	GetEditors() []Editor
}
