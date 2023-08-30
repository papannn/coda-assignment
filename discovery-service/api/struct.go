package api

type Service struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	IsActive bool   `json:"is_active"`
}

type ServiceList struct {
	Services []Service `json:"list"`
	Index    int64     `json:"index"`
}
