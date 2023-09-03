package domain

type Service struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	IsActive bool   `json:"isActive"`
}

type ServiceList struct {
	Services []*Service `json:"services"`
	Index    int64      `json:"index"`
}
