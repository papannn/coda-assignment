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
