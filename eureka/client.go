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

		serviceTTL := 30
		instance := eureka.NewInstanceInfo(hostname, e.appName, ip, e.port, uint(serviceTTL), false)

		//instance.HealthCheckUrl = ""
		//instance.HomePageUrl = ""
		instance.StatusPageUrl = applicationUri + "/actuator/info"

		instance.Metadata = &eureka.MetaData{
			Map: make(map[string]string),
		}
		instance.Metadata.Map["foo"] = "bar" //add metadata for example

		e.client.RegisterInstance(e.appName, instance) // Register new instance in your eureka(s)

		for ; ; {
			e.client.SendHeartbeat(instance.App, instance.HostName)
			time.Sleep(time.Second * time.Duration(serviceTTL-5))
		}
	}()
}

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
