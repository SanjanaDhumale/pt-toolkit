package registry

func AllServices() []Service {

	unique := make(map[string]Service)

	for _, services := range ToolRegistry {

		for _, service := range services {
			unique[service.Name] = service
		}
	}

	result := make([]Service, 0, len(unique))

	for _, service := range unique {
		result = append(result, service)
	}

	return result
}