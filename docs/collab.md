# Code Constraints

- Enums that describe a template token should start with T and end with Tokens. Example _TProjectStructureTokens_

# FileSpecDefiner

FileSpecDefiner are structs responsible for generating a FileSpec.
A file spec is a set of definition for file, this contains :

- **name** : file name
- **path** : generation path
- **content** if available, otherwise an empty string.
- **extension** : file extension (.ts for example)

A FileSpecDefiner is a struct that implement the FileSpecDefiner interface.
This include defining the following method :

```go
type FileSpecDefiner interface {
	GetFileSpec() (*FileSpec, error)
}
```

# Wrapper

Each generation process will be wrapped into an entity called Wrappers.
A wrapper is a struct that implement the following method :

```go
type Wrapper interface {
	GetDefiners() []FileSpecDefiner
}
```

The wrappers will be passes to the Generator function responsible for processing the FileSpec and taking generational actions.
The wrapper exits to abstract the the internal implementation from the CMD entry points and avoid boilerplate and intialization headaches in the entry point.
