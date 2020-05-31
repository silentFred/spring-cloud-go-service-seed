package ribbon

import (
	eureka2 "go-eureka/pkg/clients/eureka"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type NetClient struct {
	eureka     eureka2.Eureka
	httpClient *http.Client
}

func New(e eureka2.Eureka) NetClient {

	return NetClient{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		eureka: e,
	}
}

func (n NetClient) Get(u string) (resp *http.Response, err error) {
	b, _ := url.Parse(u)

	serviceName := b.Host
	path := b.Path
	app := n.eureka.GetRandomServiceInstance(serviceName)

	return n.httpClient.Get(strings.TrimRight(app.HomePageUrl, "/") + path)
}
