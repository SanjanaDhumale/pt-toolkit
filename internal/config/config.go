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

	Tools struct {

		JMeter struct {
			Enabled   bool
			Image     string
			Container string
			Version   string
		}

		K6 struct {
			Enabled bool
			Image   string
		}

		Selenium struct {
			Enabled bool
			Image   string
		}
	}
	Monitoring struct {
		Network string `yaml:"network"`

		Grafana struct {
			Enabled   bool   `yaml:"enabled"`
			Image     string `yaml:"image"`
			Container string `yaml:"container"`
			Port      string `yaml:"port"`
			Volume    string `yaml:"volume"`
		} `yaml:"grafana"`

		Prometheus struct {
			Enabled   bool   `yaml:"enabled"`
			Image     string `yaml:"image"`
			Container string `yaml:"container"`
			Port      string `yaml:"port"`
			Volume    string `yaml:"volume"`
		} `yaml:"prometheus"`

		InfluxDB struct {
			Enabled   bool   `yaml:"enabled"`
			Image     string `yaml:"image"`
			Container string `yaml:"container"`
			Port      string `yaml:"port"`
			Volume    string `yaml:"volume"`
		} `yaml:"influxdb"`
	} `yaml:"monitoring"`
}