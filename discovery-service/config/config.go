package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	HealthCheckTimeInterval int64 `json:"health_check_time_interval"`
	TimeoutThreshold        int64 `json:"timeout_threshold"`
}

var (
	ConfigObj = &Config{}
)

func initializeConfig() error {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		log.Println("error can't get the current directory")
		return err
	}

	jsonConfig, err := os.Open(fmt.Sprintf("%s%s", currentDir, "/config/config.json"))
	if err != nil {
		log.Println(err)
		log.Println("error can't find config.json file, please create it on config folder")
		return err
	}

	configByte, err := io.ReadAll(jsonConfig)
	if err != nil {
		log.Println(err)
		log.Println("error on the data io")
		return err
	}

	err = json.Unmarshal(configByte, &ConfigObj)
	if err != nil {
		log.Println(err)
		log.Println("error reading config.json file, please use a correct json form")
		return err
	}

	return nil
}

func init() {
	initializeConfig()
}
