package health_check

import (
	"github.com/papannn/coda-assignment/discovery-service/handler"
	"log"
)

func SchedulerHealthCheck(app *handler.DiscoveryService) {
	log.Println("Health check started")
	err := app.HealthCheckLogic.HealthCheck()
	if err != nil {
		log.Println(err)
	}

	log.Println("Health check ended")
}
