package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		ConcurrentRequests int    `yaml:"CONCURRENT_REQUESTS"`
		MaxRequests        int    `yaml:"MAX_REQUESTS"`
		Port               string `yaml:"PORT"`
	} `yaml:"server"`
	Picture struct {
		StartDate        string `yaml:"START_DATE"`
		EndDate           string `yaml:"END_DATE"`
	} `yaml:"picture"`
	External struct {
		ApiKey        string `yaml:"API_KEY"`
		Url           string `yaml:"URL"`
		DateParameter string `yaml:"DATE_PARAMETER"`
	} `yaml:"external"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config) {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
