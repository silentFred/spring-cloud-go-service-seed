package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jamowei/senv"
	"go-eureka/pkg/clients/eureka"
	"go-eureka/pkg/clients/flights"
	"strconv"
)

func main() {

	appName := "go-fish-api"

	config := senv.NewConfig(
		"localhost",
		"9000",
		appName,
		[]string{"test"},
		"master")

	config.Fetch(true, true)
	config.Process()

	applicationPort, _ := strconv.Atoi(config.Properties["server.port"])

	eurekaHosts := []string{
		config.Properties["eureka.client.service-url.default-zone"],
	}

	eurekaClient := eureka.New(appName, applicationPort, eurekaHosts)
	flightsClient := flights.New(eurekaClient)
	app := gin.Default()

	app.GET("/flights/:id", func(context *gin.Context) {
		flightIdParameter := context.Params.ByName("id")
		flightId, _ := strconv.Atoi(flightIdParameter)
		flightResponse, _ := flightsClient.GetFlight(flightId)
		context.JSON(200, flightResponse)
	})

	app.GET("/actuator/env", func(context *gin.Context) {
		context.JSONP(200, config.Properties)
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
