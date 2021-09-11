package config

import (
	"github.com/BurntSushi/toml"
	"github.com/subchen/go-log"
	"os"
	"path/filepath"
)

const configName = "server.cfg"

// Current configuration
var config *serverConfig

// Init - A method that reads the configuration file and initializes it in the service
func Init() {
	ex, err := os.Executable()
	if err != nil {
		log.Panicf("error: directory not found or does not exist: %v", err)
	}

	file := filepath.Join(filepath.Dir(ex), configName)

	var cfg serverConfig
	_, err = toml.DecodeFile(file, &cfg)
	if err != nil {
		log.Panicf("error: could not read the config file: %v", err)
	}

	config = &cfg

	log.Infof("config file initialized: %s", file)
	log.Debugf("config: %v", config)
}

// Get - returns the current configuration
func Get() *serverConfig {
	return config
}

type serverConfig struct {
	Port int
	Log  logConfig
}

type logConfig struct {
	FileName     string
	MaxCountFile int
	Level        string
}
