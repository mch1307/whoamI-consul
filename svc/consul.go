package svc

import (
	"fmt"
	"strconv"

	api "github.com/hashicorp/consul/api"
)

//NewClient returns a consul client instance
func NewClient(consulAddr, consulPort, consulToken string) (Consul *api.Client) {
	config := api.DefaultConfig()
	config.Address = consulAddr + ":" + consulPort
	config.Token = consulToken
	consul, err := api.NewClient(config)
	if err != nil {
		fmt.Println("error connecting to consul ", err)
	}
	return consul
}

// RegisterService register the service to Consul service catalog
func RegisterService(ag *api.Agent, service, hostname, protocol string, port int) error {

	serviceID := service
	consulService := api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    service,
		Tags:    []string{service},
		Port:    port,
		Address: hostname,
		Check: &api.AgentServiceCheck{
			Script:   "curl --connect-timeout=5 " + protocol + "://" + hostname + ":" + strconv.Itoa(port),
			Interval: "10s",
			Timeout:  "8s",
			TTL:      "",
			HTTP:     protocol + "://" + hostname + ":" + strconv.Itoa(port) + "/health",
			Status:   "passing",
		},
		Checks: api.AgentServiceChecks{},
	}
	err := ag.ServiceRegister(&consulService)
	if err != nil {
		return err
	}

	return err
}

// DeregisterService unregister the service from Consul service catalog
func DeregisterService(ag *api.Agent, service string) {
	_ = ag.ServiceDeregister(service)

}
