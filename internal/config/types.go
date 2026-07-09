package config

type Tool struct {
	Enabled   bool   `yaml:"enabled"`
	Image     string `yaml:"image"`
	Container string `yaml:"container,omitempty"`
	Version   string `yaml:"version,omitempty"`
}

type Service struct {
	Enabled   bool   `yaml:"enabled"`
	Image     string `yaml:"image"`
	Container string `yaml:"container"`
	Port      string `yaml:"port"`
	Volume    string `yaml:"volume"`
}
