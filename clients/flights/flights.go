package flights

import (
	"fmt"
	eureka2 "go-eureka/clients/eureka"
	"io/ioutil"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

type Flight struct {
	eureka eureka2.Eureka
}

func New(e eureka2.Eureka) Flight {
	return Flight{eureka: e}
}

func (f Flight) GetFlight(flightId int) []byte {
	serviceName := "FLIGHT-SERVICE"
	app := f.eureka.GetRandomServiceInstance(serviceName)

	rawResponse, _ := netClient.Get(fmt.Sprintf("%sflights/%d", app.HomePageUrl, flightId))
	response, _ := ioutil.ReadAll(rawResponse.Body)
	defer rawResponse.Body.Close()
	return response
}
