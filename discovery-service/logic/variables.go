package logic

type Service struct {
	IP   string
	Port string
}

type ServiceList struct {
	Services []Service
	Index    int64
}

var (
	ServiceMap map[string]*ServiceList
)

func init() {
	ServiceMap = make(map[string]*ServiceList)
}
