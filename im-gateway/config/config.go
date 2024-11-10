package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type TCPServerConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	WorkerNum int    `yaml:"worker-num"`
}

type RPCServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type EtcdConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AppConfig struct {
	Debug                   bool   `yaml:"debug"`
	Name                    string `yaml:"name"`
	HeartbeatIntervalMs     int    `yaml:"heartbeat-interval-ms"`
	CloseConnIntervalMs     int    `yaml:"close-conn-interval-ms"`
	MetricsUpdateIntervalMs int    `yaml:"metrics-update-interval-ms"`
	TCPAdvertiseAddrport    string `yaml:"tcp-advertise-addrport"`
	RPCAdvertiseAddrport    string `yaml:"rpc-advertise-addrport"`
}

type Config struct {
	TCPServerConfig TCPServerConfig `yaml:"tcp-server"`
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
