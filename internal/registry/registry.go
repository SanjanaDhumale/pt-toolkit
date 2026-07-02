package registry

type Service struct {
    Name        string
    Container   string
    Image       string
    Description string
    Enabled     bool
    Ports       []string
}

var ToolRegistry = map[string][]Service{

	"jmeter": {
		{
			Name:        "JMeter",
			Container:   "pt-jmeter",
			Image:       "pt-jmeter:v1",
			Description: "Apache JMeter",
		},
		{
			Name:        "Grafana",
			Container:   "pt-grafana",
			Image:       "grafana/grafana:latest",
			Description: "Monitoring Dashboard",
		},
		{
			Name:        "Prometheus",
			Container:   "pt-prometheus",
			Image:       "prom/prometheus:latest",
			Description: "Metrics Collector",
		},
		{
			Name:        "InfluxDB",
			Container:   "pt-influxdb",
			Image:       "influxdb:2.7",
			Description: "Time Series Database",
		},
	},

	"k6": {
		{
			Name:        "K6",
			Container:   "pt-k6",
			Image:       "grafana/k6:latest",
			Description: "Grafana K6",
		},
		{
			Name:        "Grafana",
			Container:   "pt-grafana",
			Image:       "grafana/grafana:latest",
			Description: "Monitoring Dashboard",
		},
	},

	"selenium": {
		{
			Name:        "Selenium",
			Container:   "pt-selenium",
			Image:       "selenium/standalone-chrome:latest",
			Description: "Browser Automation",
		},
	},
}