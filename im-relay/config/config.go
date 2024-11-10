package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type RPCServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type EtcdConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AppConfig struct {
	Debug                bool   `yaml:"debug"`
	RPCAdvertiseAddrport string `yaml:"rpc-advertise-addrport"`
}

type Config struct {
	RPCServerConfig RPCServerConfig `yaml:"rpc-server"`
	EtcdConfig      []EtcdConfig    `yaml:"etcd"`
	AppConfig       AppConfig       `yaml:"app"`
}

func MustGetConfigFromFile(file string) *Config {
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	obj := Config{}
	err = yaml.Unmarshal(bs, &obj)
	if err != nil {
		panic(err)
	}
	return &obj
}
