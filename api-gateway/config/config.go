package config

type Config struct {
	IP                      string `json:"ip"`
	Port                    string `json:"port"`
	DiscoveryServiceBaseUrl string `json:"discovery_service_base_url"`
}
