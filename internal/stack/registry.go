package stack

var Registry = map[string]Stack{

	"protocol": {

		Name:        "Protocol Testing",
		Description: "JMeter based performance testing",
		Version:     "1.0",

		Components: []Component{

			{
				Name:      "Apache JMeter",
				Type:      "Docker",
				Version:   "5.6.3",
				Image:     "pt-jmeter:v1",
				Container: "pt-jmeter",
				Required:  true,
			},

			{
				Name:      "Java",
				Type:      "System",
				Version:   "21",
				Required:  true,
			},
		},
	},

	"monitoring": {

		Name:        "Monitoring",
		Description: "Monitoring Dashboard",
		Version:     "1.0",

		Components: []Component{

			{
				Name:      "Grafana",
				Type:      "Docker",
				Version:   "latest",
				Image:     "grafana/grafana:latest",
				Container: "pt-grafana",
				Required:  true,
			},

			{
				Name:      "Prometheus",
				Type:      "Docker",
				Version:   "latest",
				Image:     "prom/prometheus:latest",
				Container: "pt-prometheus",
				Required:  true,
			},

			{
				Name:      "InfluxDB",
				Type:      "Docker",
				Version:   "2.7",
				Image:     "influxdb:2.7",
				Container: "pt-influxdb",
				Required:  true,
			},
		},
	},
}