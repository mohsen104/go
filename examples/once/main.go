package main

import "sync"

type Config struct {
	ConnectionString string
}

var (
	once   sync.Once
	config *Config
)

func main() {
	for i := 0; i < 10; i++ {
		cfg := GetConfig()
		println(i, " : ", &cfg.ConnectionString)
	}
}

func GetConfig() *Config {
	once.Do(func() {
		println("Creating Config")
		config = &Config{}
	})
	return config
}
