package main

import (
	"./api"
	"./config"
	consulcli "./consul"
	"./osinfo"
	serf "./serf"
	"fmt"
	flag "github.com/spf13/pflag"
	"time"
)

func errorHandle(err error) error {
	if err != nil {
		fmt.Errorf("An error has occured %g", err)
		panic(err)
	}
	return nil
}

func main() {

	osinfo.StartTime = time.Now()
	osinfo.IPAddress, osinfo.Hostname = osinfo.GetLocalIP()
	osinfo.TotalMem, osinfo.FreeMem, osinfo.UsedMem = osinfo.GetMemInfo()

	// parse command line for config file option or use default
	var configfile string
	flag.StringVarP(&configfile, "config", "c", "", "Path to configuration file")
	flag.Parse()

	if configfile == "" {
		configfile = "/Users/kabae/go_workspace/dsmprov/config.toml"
	}

	cfg := config.ReadConfig(configfile)

	// connect to Consul server;
	cons, err := consulcli.NewConsulClient(cfg.Discovery.Server)
	errorHandle(err)
	// cons.Register(osinfo.Hostname, osinfo.IPAddress, 5050)

	serfcli, err := serf.NewSerfClient(cfg.Discovery.Server)
	errorHandle(err)

	go api.Start(cfg.Api, cons, serfcli)
	<-(chan string)(nil)

}
