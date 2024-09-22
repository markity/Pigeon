package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

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

type Config struct {
	MysqlConfig MysqlConfig `yaml:"mysql"`
	RedisConfig RedisConfig `yaml:"redis"`
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
