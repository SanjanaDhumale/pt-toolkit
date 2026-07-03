package docker

import "fmt"

func StartMonitoringStack(network string) error {

	CreateVolume("pt-grafana-data")
	CreateVolume("pt-prometheus-data")
	CreateVolume("pt-influxdb-data")

	if err := RunContainer(
		"pt-influxdb",
		"influxdb:2.7",
		network,
		"8086:8086",
	); err != nil {
		return err
	}

	if err := RunContainer(
		"pt-prometheus",
		"prom/prometheus:latest",
		network,
		"9090:9090",
	); err != nil {
		return err
	}

	if err := RunContainer(
		"pt-grafana",
		"grafana/grafana:latest",
		network,
		"3000:3000",
	); err != nil {
		return err
	}

	if !IsContainerRunning("pt-grafana") {
		return fmt.Errorf("Grafana failed to start")
	}

	if !IsContainerRunning("pt-prometheus") {
		return fmt.Errorf("Prometheus failed to start")
	}

	if !IsContainerRunning("pt-influxdb") {
		return fmt.Errorf("InfluxDB failed to start")
	}

	return nil
}