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

type MysqlConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	Db   string `yaml:"db"`
}

type RedisConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	KeyPrefix string `yaml:"keyprefix"`
}

type AppConfig struct {
	RPCAdvertiseAddrport string `yaml:"rpc-advertise-addrport"`
	DeviceNumLimit       int    `yaml:"device-num-limit"`
}

type Config struct {
	RPCServerConfig RPCServerConfig `yaml:"rpc-server"`
	EtcdConfig      []EtcdConfig    `yaml:"etcd"`
	MysqlConfig     MysqlConfig     `yaml:"mysql"`
	RedisConfig     RedisConfig     `yaml:"redis"`
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
