package config

import (
	"os"
)

var logger = log.logger.New().Name("config")

func Load(config interface{}) {
	configStr := os.Getenv("CONFIG")
	if configStr == "" {
		log.Info("Config env not set. Using default config")
		return
	}

	err := json.Unmarshal([]byte(configStr), config)
	if err != nil {
		//use zap go.uber.org/zap
		log.Fatal("error loading config json")
	}
}
