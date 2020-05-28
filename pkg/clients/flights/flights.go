package flights

import (
	"fmt"
	eureka2 "go-eureka/pkg/clients/eureka"
	"go-eureka/pkg/clients/ribbon"
	"io/ioutil"
	"net/http"
)

type GORibbon interface {
	Get(url string) (resp *http.Response, err error)
}

type Flight struct {
	ribbonClient GORibbon
}

func New(e eureka2.Eureka) Flight {
	return Flight{ribbonClient: ribbon.New(e)}
}

func (f Flight) GetFlight(flightId int) []byte {
	rawResponse, _ := f.ribbonClient.Get(fmt.Sprintf("http://flight-service/flights/%d", flightId))
	response, _ := ioutil.ReadAll(rawResponse.Body)
	defer rawResponse.Body.Close()
	return response
}
