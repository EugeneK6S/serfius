package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Api       ApiConfig
	Discovery DiscoveryConfig
}

type ApiConfig struct {
	Bind string `toml:"bind"`
}

type DiscoveryConfig struct {
	Server string `toml:"server"`
}

func ReadConfig(configfile string) Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
