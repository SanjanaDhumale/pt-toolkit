package stack

import "github.com/SanjanaDhumale/pt-toolkit/internal/component"

var Registry = map[string]Stack{

	"protocol": {
		Name: "Protocol Testing",
		Components: []component.Component{
			component.Java{},
		},
	},

	"monitoring": {
		Name: "Monitoring",
		Components: []component.Component{},
	},
}