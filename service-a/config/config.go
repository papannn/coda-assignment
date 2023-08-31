package config

type Config struct {
	Namespace               string `json:"namespace"`
	IP                      string `json:"ip"`
	Port                    string `json:"port"`
	DiscoveryServiceBaseURL string `json:"discovery_service_base_url"`
}
