package types

type Config struct {
	Template string `json:"template"`
	Style    string `json:"style"`
	Dist     string `json:"dist"`
	Type     string `json:"type"`
}
