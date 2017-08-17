package svc

import (
	"fmt"
	"strconv"

	api "github.com/hashicorp/consul/api"
)

// NewClient returns a consul client instance
func NewClient(consulAddr, consulPort, consulToken string) (Consul *api.Client, err error) {
	config := api.DefaultConfig()
	config.Address = consulAddr + ":" + consulPort
	config.Token = consulToken
	consul, err := api.NewClient(config)
	if err != nil {
		fmt.Println("error connecting to consul ", err)
	}
	return consul, err
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
			HTTP:     protocol + "://" + "172.30.16.47" + ":" + strconv.Itoa(port) + "/health",
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

// GetKV get value of key from Consul K/V store and returns as string
func GetKV(cli *api.Client, key string) (string, error) {
	var ret string
	kv := cli.KV()
	val, _, err := kv.Get(key, nil)
	if err != nil {
		fmt.Println(err)
		val.Value = nil
	}
	if val != nil {
		ret = string(val.Value)
	} else {
		ret = "not_found"
	}
	return ret, err
}
