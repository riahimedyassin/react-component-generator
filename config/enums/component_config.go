package enums

type Styling string

const (
	NONE     Styling = "none"
	SCSS     Styling = "scss"
	CSS      Styling = "css"
	TAILWIND Styling = "tailwind"
)

type ComponentType string

const (
	FUNCTIONAL ComponentType = "functional"
	CLASS      ComponentType = "class"
)

type ProjectBase string

const (
	TYPESCRIPT ProjectBase = "typescript"
	JAVASCRIPT ProjectBase = "javascript"
)
