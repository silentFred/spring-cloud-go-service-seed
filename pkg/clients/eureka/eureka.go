package eureka

import (
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"math/rand"
	"net"
	"os"
	"time"
)

func New(appName string, port int, hosts []string) Eureka {
	var client = eureka.NewClient(hosts)

	return Eureka{
		appName: appName,
		port:    port,
		hosts:   hosts,
		client:  client,
	}
}

type Eureka struct {
	appName string
	port    int
	hosts   []string
	client  *eureka.Client
}

func (e Eureka) RegisterService() {
	go func() {
		hostname, _ := os.Hostname()
		applicationUri := fmt.Sprintf("http://%s:%d", hostname, e.port)
		ip := getIp()

		serviceTTL := 30 // todo configure
		instance := eureka.NewInstanceInfo(hostname, e.appName, ip, e.port, uint(serviceTTL), false)

		instance.StatusPageUrl = applicationUri + "/actuator/info"
		instance.HealthCheckUrl = applicationUri + "/actuator/health"

		instance.Metadata = &eureka.MetaData{
			Map: make(map[string]string),
		}

		instance.Metadata.Map["service-description"] = "GoLang service running on Spring Cloud"

		e.client.RegisterInstance(e.appName, instance)

		for {
			e.client.SendHeartbeat(instance.App, instance.HostName)
			time.Sleep(time.Second * time.Duration(serviceTTL-5))
		}
	}()
}

// TODO implement local service instance cache that updates every x seconds
func (e Eureka) GetRandomServiceInstance(s string) eureka.InstanceInfo {
	flightServices, _ := e.client.GetApplication(s)
	apps := flightServices.Instances
	randomIndex := rand.Intn(len(apps))
	app := apps[randomIndex]
	return app
}

func getIp() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}
