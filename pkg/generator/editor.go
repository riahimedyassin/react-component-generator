package generator

type Editor interface {
	GetEditSpecs() (*EditFileSpec, error)
}

type EditFileSpec struct {
	Name       string
	Path       string
	Extension  string
	NewContent string
}

func NewEditFileSpec(
	Name string,
	Path string,
	Extension string,
	NewContent string,
) *EditFileSpec {
	return &EditFileSpec{
		Name:       Name,
		Path:       Path,
		Extension:  Extension,
		NewContent: NewContent,
	}
}
