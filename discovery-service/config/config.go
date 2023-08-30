package config

type Config struct {
	IP                      string `json:"ip"`
	Port                    string `json:"port"`
	HealthCheckTimeInterval int64  `json:"health_check_time_interval"`
	TimeoutThreshold        int64  `json:"timeout_threshold"`
	LoadBalancingAlgorithm  string `json:"load_balancing_algorithm"`
}
