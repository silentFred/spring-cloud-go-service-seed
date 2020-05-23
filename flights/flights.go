package flights

import (
	"fmt"
	"go-eureka/eureka"
	"io/ioutil"
	"net/http"
)

type Flight struct {
	eureka eureka.Eureka
}

func New(e eureka.Eureka) Flight {
	return Flight{eureka: e}
}

func (f Flight) GetFlight(flightId int) []byte {
	serviceName := "FLIGHT-SERVICE"
	app := f.eureka.GetRandomServiceInstance(serviceName)
	rawResponse, _ := http.Get(fmt.Sprintf("%sflights/%d", app.HomePageUrl, flightId))
	response, _ := ioutil.ReadAll(rawResponse.Body)
	defer rawResponse.Body.Close()
	return response
}
