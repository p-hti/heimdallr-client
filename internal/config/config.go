package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Broker Kafka `yaml:"broker"`
}

type Kafka struct {
	KafkaAddress string `yaml:"address"`
	KafkaTopic   string `yaml:"topic"`
}

const configPath = "config/config.yaml"

func MustLoad() *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found!")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
