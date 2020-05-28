package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-eureka/pkg/clients/eureka"
	"go-eureka/pkg/clients/flights"
)

func main() {

	appName := "go-fish-api"

	applicationPort := 9999

	eurekaHosts := []string{
		"http://localhost:8761/eureka",
	}

	eurekaClient := eureka.New(appName, applicationPort, eurekaHosts)

	flightsClient := flights.New(eurekaClient)

	app := gin.Default()

	app.GET("/actuator/info", func(context *gin.Context) {
		context.JSONP(200, string(flightsClient.GetFlight(1)))
	})

	app.GET("/actuator/health", func(context *gin.Context) {
		context.JSONP(200, "OK")
	})

	eurekaClient.RegisterService()

	app.Run(fmt.Sprintf(":%d", applicationPort))
}
