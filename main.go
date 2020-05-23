package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go-eureka/eureka"
	"go-eureka/flights"
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

	eurekaClient.RegisterService()

	app.Run(fmt.Sprintf(":%d", applicationPort))
}
