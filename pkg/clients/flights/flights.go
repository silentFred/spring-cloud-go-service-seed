package flights

import (
	"encoding/json"
	"fmt"
	eureka2 "go-eureka/pkg/clients/eureka"
	"go-eureka/pkg/clients/ribbon"
	"io/ioutil"
	"net/http"
)

type FlightResponse struct {
	Id              int     `json:"id"`
	FlightDate      string  `json:"flightDate"`
	Origin          string  `json:"origin"`
	Destination     string  `json:"destination"`
	SeatsAvailable  int     `json:"seatsAvailable"`
	Price           float32 `json:"price"`
	CostServicePort int     `json:"costServicePort"`
}

type GORibbon interface {
	Get(url string) (resp *http.Response, err error)
}

type Flight struct {
	ribbonClient GORibbon
}

func New(e eureka2.Eureka) Flight {
	return Flight{ribbonClient: ribbon.New(e)}
}

func (f Flight) GetFlight(flightId int) (FlightResponse, error) {
	rawResponse, _ := f.ribbonClient.Get(fmt.Sprintf("http://flight-service/flights/%d", flightId))
	response, _ := ioutil.ReadAll(rawResponse.Body)
	defer rawResponse.Body.Close()

	flightResponse := FlightResponse{}

	err := json.Unmarshal(response, &flightResponse)

	if err != nil {
		return FlightResponse{}, err
	}

	return flightResponse, err
}
