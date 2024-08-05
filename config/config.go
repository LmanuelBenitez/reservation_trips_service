package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	"template_soa/libs/logger"
	"template_soa/models"

	"github.com/kelseyhightower/envconfig"
)

var (
	_, path, _, _ = runtime.Caller(0)
	basePath      = filepath.Dir(path)
	enviroment    = "dev"
	configuration = models.Configuration{}
)

func NewConfiguration() *models.Configuration {
	validateEnvironment()
	readConfigFile()
	readConfigEnv()

	return &configuration
}

func validateEnvironment() {
	if environmentValue, exists := os.LookupEnv("GO_ENV"); exists {
		enviroment = environmentValue
	}
}

func readConfigFile() {
	directory := basePath + "/config" + enviroment + ".yml"

	file, err := os.Open(directory)
	if err != nil {
		logger.Fatal("Configuration", "readConfigFile", err.Error())
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(configuration); err != nil {
		logger.Fatal("Configuration", "readConfigFile", err.Error())
	}
}

func readConfigEnv() {
	if err := envconfig.Process("", &configuration); err != nil {
		logger.Fatal("Configuration", "readConfigEnv", err.Error())
	}
}
