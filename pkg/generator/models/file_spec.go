package generator_models

type FileSpec struct {
	Name      string
	Content   string
	Path      string
	Extension string
}

func NewFileSpec(
	Name string,
	Path string,
	Extension string,
	Content string,
) *FileSpec {
	return &FileSpec{
		Name:      Name,
		Content:   Content,
		Path:      Path,
		Extension: Extension,
	}
}
