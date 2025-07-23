package config

type coreConfig struct {
	Structure string `json:"structure"`
	Template  string `json:"template"`
	Styling   string `json:"styling"`
}

type stateConfig struct {
	Type string `json:"type"`
}

type lintingConfig struct {
	Eslint   bool `json:"eslint"`
	Prettier bool `json:"prettier"`
}

type GlobalConfig struct {
	Core    coreConfig    `json:"core"`
	State   stateConfig   `json:"state"`
	Linting lintingConfig `json:"linting"`
}
