package registry

import "fmt"

func Lookup(tool string) ([]Service, error) {

	services, ok := ToolRegistry[tool]

	if !ok {
		return nil, fmt.Errorf("unknown tool")
	}

	return services, nil
}