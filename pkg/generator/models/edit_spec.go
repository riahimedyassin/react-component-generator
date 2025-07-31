package generator_models

// defines editable files.
type EditFileSpec struct {
	Name           string
	Path           string
	Extension      string
	NewContent     string
	CurrentContent string
}

func NewEditFileSpec(
	Name string,
	Path string,
	Extension string,
	NewContent string,
	CurrentContent string,
) *EditFileSpec {
	return &EditFileSpec{
		Name:           Name,
		Path:           Path,
		Extension:      Extension,
		NewContent:     NewContent,
		CurrentContent: CurrentContent,
	}
}
