package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HTTPServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type EtcdConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AppConfig struct {
	AdvertiseAddrport string `yaml:"tcp-advertise-addrport"`
}

type Config struct {
	HTTPServerConfig HTTPServerConfig `yaml:"http-server"`
	EtcdConfig       []EtcdConfig     `yaml:"etcd"`
	AppConfig        AppConfig        `yaml:"app"`
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
