package config

type Config struct {
	Toolkit struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"toolkit"`

	Docker struct {
		Image       string `yaml:"image"`
		Dockerfile  string `yaml:"dockerfile"`
		ComposeFile string `yaml:"compose_file"`
	} `yaml:"docker"`

	Workspace struct {
		Reports string `yaml:"reports"`
		Logs    string `yaml:"logs"`
	} `yaml:"workspace"`

	Plugins struct {
		JMeter bool `yaml:"jmeter"`
		K6      bool `yaml:"k6"`
	} `yaml:"plugins"`
}