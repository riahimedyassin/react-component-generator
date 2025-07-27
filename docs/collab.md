# Code Constraints

- Enums that describe a template token should start with T and end with Tokens. Example _TProjectStructureTokens_

# Generators

Generators are structs responsible for generating a FileSpec.
A file spec is a set of definition for file, this contains :

- **name** : file name
- **path** : generation path
- **content** if available, otherwise an empty string.
- **extension** : file extension (.ts for example)

A generator is a struct that implement the FileSpecDefiner interface.
This include defining the following method :

```go
GetFileSpec() (*FileSpec, error)
```
