package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-eureka/pkg/clients/eureka"
	"go-eureka/pkg/clients/flights"
	"strconv"
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

	app.GET("/flights/:id", func(context *gin.Context) {
		fligthId := context.Params.ByName("id")
		id, _ := strconv.Atoi(fligthId)
		flight, _ := flightsClient.GetFlight(id)
		context.JSON(200, flight)
	})

	app.GET("/actuator/info", func(context *gin.Context) {
		context.JSONP(200, "Some info here")
	})

	app.GET("/actuator/health", func(context *gin.Context) {
		context.JSONP(200, "OK")
	})

	eurekaClient.RegisterService()

	app.Run(fmt.Sprintf(":%d", applicationPort))
}
