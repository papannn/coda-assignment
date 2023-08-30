package logic

type Service struct {
	IP       string
	Port     string
	IsActive bool
}

type ServiceList struct {
	Services []Service
	Index    int64
}

type ServiceMap = map[string]*ServiceList
