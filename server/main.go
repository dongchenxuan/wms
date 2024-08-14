package main

import (
	"flag"
	"wms/conf"
	"wms/log"
	"wms/router"
)

func main() {
	var configFile string
	var err error

	flag.StringVar(&configFile, "conf", "./conf/conf_dev.json", "config file")
	flag.Parse()

	if configFile == "" {
		log.Errorf("config file argument is required.")
		return
	}

	err = conf.Init(configFile)
	if err != nil {
		log.Errorf("[init][conf] %v", err.Error())
		return
	}

	router.Init()
}
