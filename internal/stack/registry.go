package stack

import "github.com/SanjanaDhumale/pt-toolkit/internal/component"

var Registry = map[string]Stack{

	"protocol": {
		Name: "Protocol Testing",
		Network: "pt-protocol-network",
		Components: []component.Component{
			component.Java{},
			component.Docker{},
			component.JMeter{},
		},
	},

	"monitoring": {
		Name: "Monitoring",
		Network: "pt-monitoring-network",
		Components: []component.Component{
			component.Grafana{},
			component.Prometheus{},
			component.InfluxDB{},
		},
	},

	"browser": {
		Name: "Browser Automation",
		Network: "pt-browser-network",
		Components: []component.Component{
			component.Selenium{},
		},
	},
}