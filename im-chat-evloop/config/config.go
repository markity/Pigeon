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

type AppConfig struct {
	Debug bool `yaml:"debug"`
	// 供snowflake使用
	NodeId               int64  `yaml:"node-id"`
	RPCAdvertiseAddrport string `yaml:"rpc-advertise-addrport"`
}

type Config struct {
	RPCServerConfig RPCServerConfig `yaml:"rpc-server"`
	EtcdConfig      []EtcdConfig    `yaml:"etcd"`
	MysqlConfig     MysqlConfig     `yaml:"mysql"`
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
