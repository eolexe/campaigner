package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Name   string       `json:"name"`
	Server HostPortPair `json:"server"`
}

// HostPortPair is a pair of host address and a port.
type HostPortPair struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (h HostPortPair) String() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func MustNewConfig(path string) Config {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(fmt.Sprintf("Can't load config file using provided path. Details: %s", err.Error()))
	}

	conf := Config{}
	err = json.Unmarshal(data, &conf)

	if err != nil {
		panic(fmt.Sprintf("Can't parse the content of the config file. Details: %s", err.Error()))
	}

	return conf
}
