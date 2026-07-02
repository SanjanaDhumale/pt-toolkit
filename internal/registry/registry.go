package registry

type Service struct {
	Name        string
	Container   string
	Description string
}

var ToolRegistry = map[string][]Service{

	"jmeter": {
		{
			Name:        "JMeter",
			Container:   "pt-jmeter",
			Description: "Apache JMeter",
		},
		{
			Name:        "Grafana",
			Container:   "pt-grafana",
			Description: "Monitoring Dashboard",
		},
		{
			Name:        "Prometheus",
			Container:   "pt-prometheus",
			Description: "Metrics Collector",
		},
		{
			Name:        "InfluxDB",
			Container:   "pt-influxdb",
			Description: "Time Series Database",
		},
	},

	"k6": {
		{
			Name:        "K6",
			Container:   "pt-k6",
			Description: "Grafana K6",
		},
		{
			Name:        "Grafana",
			Container:   "pt-grafana",
			Description: "Monitoring Dashboard",
		},
	},

	"selenium": {
		{
			Name:        "Selenium",
			Container:   "pt-selenium",
			Description: "Browser Automation",
		},
	},
}