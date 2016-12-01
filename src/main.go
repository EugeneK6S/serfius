package main

import (
	"./api"
	"./config"
	// consulcli "./consul"
	"./osinfo"
	// serf "./serf"
	// "fmt"
	flag "github.com/spf13/pflag"
	"time"
)

func main() {

	osinfo.StartTime = time.Now()
	osinfo.IPAddress, osinfo.Hostname = osinfo.GetLocalIP()
	osinfo.TotalMem, osinfo.FreeMem, osinfo.UsedMem = osinfo.GetMemInfo()

	// parse command line for config file option or use default
	var configfile string
	flag.StringVarP(&configfile, "config", "c", "", "Path to configuration file")
	flag.Parse()

	// Use default config location if not specified
	if configfile == "" {
		configfile = "./config/config.toml"
	}

	cfg, _ := config.ReadConfig(configfile)

	go api.Start(cfg)
	<-(chan string)(nil)

}
