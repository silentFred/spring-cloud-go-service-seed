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

	e := eureka.New(appName, applicationPort, eurekaHosts)

	f := flights.New(e)

	app := gin.Default()

	app.GET("/actuator/info", func(context *gin.Context) {
		context.JSONP(200, string(f.GetFlight(1)))
	})

	e.RegisterService()
	app.Run(fmt.Sprintf(":%d", applicationPort))
}

