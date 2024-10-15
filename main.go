package main

import (
	"fmt"

	yamlconfig "github.com/maeck70/yaml-config"
)

type RabbitMQ_t struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

type myConfig_t struct {
	Name     string                `yaml:"Name"`
	City     string                `yaml:"City"`
	State    string                `yaml:"State"`
	Id       int                   `yaml:"Id"`
	Options  []string              `yaml:"Options"`
	Rabbitmq map[string]RabbitMQ_t `yaml:"Rabbitmq"`
}

func main() {
	c := yamlconfig.LoadConfig("example.config.yaml", &myConfig_t{})

	newConf := c.(*myConfig_t)

	fmt.Printf("new conf: %+v\n", newConf)
	fmt.Printf("Name: %s\n", newConf.Name)
	fmt.Printf("Rabbitmq main vhost: %s\n", newConf.Rabbitmq["main"].Vhost)
	fmt.Printf("Rabbitmq baskets vhost: %s\n", newConf.Rabbitmq["baskets"].Vhost)
}
