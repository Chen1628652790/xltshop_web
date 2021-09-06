package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	healthCheck := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	serviceRegistration := new(api.AgentServiceRegistration)
	serviceRegistration.Name = name
	serviceRegistration.ID = id
	serviceRegistration.Port = port
	serviceRegistration.Tags = tags
	serviceRegistration.Address = address
	serviceRegistration.Check = healthCheck
	if err = client.Agent().ServiceRegister(serviceRegistration); err != nil {
		zap.S().Errorw("client.Agent().ServiceRegister failed", "msg", err.Error())
		return err
	}
	return nil
}

func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	return client.Agent().ServiceDeregister(serviceId)
}
