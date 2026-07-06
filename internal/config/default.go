package config

func Default() Config {

	var cfg Config

	cfg.Toolkit.Name = "PT Toolkit"
	cfg.Toolkit.Version = "1.0.0"

	cfg.Docker.Image = "pt-toolkit"
	cfg.Docker.Dockerfile = "Dockerfile"
	cfg.Docker.ComposeFile = "docker-compose.yml"

	cfg.Workspace.Reports = "reports"
	cfg.Workspace.Logs = "logs"

	// -------------------------
	// Protocol Testing
	// -------------------------

	cfg.Tools.JMeter.Enabled = true
	cfg.Tools.JMeter.Image = "pt-jmeter-enterprise:v1"
	cfg.Tools.JMeter.Container = "pt-jmeter"
	cfg.Tools.JMeter.Version = "5.6.3"

	cfg.Tools.K6.Enabled = true
	cfg.Tools.K6.Image = "grafana/k6:latest"

	cfg.Tools.Selenium.Enabled = true
	cfg.Tools.Selenium.Image = "selenium/standalone-chrome:latest"

	// -------------------------
	// Monitoring
	// -------------------------

	cfg.Monitoring.Network = "pt-monitoring-network"

	cfg.Monitoring.Grafana.Enabled = true
	cfg.Monitoring.Grafana.Image = "grafana/grafana:latest"
	cfg.Monitoring.Grafana.Container = "pt-grafana"
	cfg.Monitoring.Grafana.Port = "3000:3000"
	cfg.Monitoring.Grafana.Volume = "pt-grafana-data"

	cfg.Monitoring.Prometheus.Enabled = true
	cfg.Monitoring.Prometheus.Image = "prom/prometheus:latest"
	cfg.Monitoring.Prometheus.Container = "pt-prometheus"
	cfg.Monitoring.Prometheus.Port = "9090:9090"
	cfg.Monitoring.Prometheus.Volume = "pt-prometheus-data"

	cfg.Monitoring.InfluxDB.Enabled = true
	cfg.Monitoring.InfluxDB.Image = "influxdb:2.7"
	cfg.Monitoring.InfluxDB.Container = "pt-influxdb"
	cfg.Monitoring.InfluxDB.Port = "8086:8086"
	cfg.Monitoring.InfluxDB.Volume = "pt-influxdb-data"

	return cfg
}