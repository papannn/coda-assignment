package api

type RegisterRequest struct {
	Namespace string `json:"namespace""`
	IP        string `json:"ip"`
	Port      string `json:"port"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type UnregisterRequest struct {
	Namespace string `json:"namespace""`
	IP        string `json:"ip"`
	Port      string `json:"port"`
}

type UnregisterResponse struct {
	Message string `json:"message"`
}

type LookupRequest struct {
	Namespace string `json:"namespace"`
}

type LookupResponse struct {
	IP                    string `json:"ip"`
	Port                  string `json:"port"`
	ServiceAvailableCount int64  `json:"service_available_count"`
}

type StatusResponse struct {
	ServiceMap map[string]ServiceList `json:"Service"`
}
