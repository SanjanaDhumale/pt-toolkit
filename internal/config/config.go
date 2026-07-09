package config

type Config struct {
	Toolkit struct {
		Name        string `yaml:"name"`
		Version     string `yaml:"version"`
		Description string `yaml:"description"`
		Author      string `yaml:"author"`
		Company     string `yaml:"company"`
		Website     string `yaml:"website"`
		License     string `yaml:"license"`
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

	Tools struct {
		JMeter   Tool `yaml:"jmeter"`
		K6       Tool `yaml:"k6"`
		Selenium Tool `yaml:"selenium"`
	} `yaml:"tools"`

	Monitoring struct {
		Network string `yaml:"network"`

		Grafana    Service `yaml:"grafana"`
		Prometheus Service `yaml:"prometheus"`
		InfluxDB   Service `yaml:"influxdb"`
	} `yaml:"monitoring"`
}
