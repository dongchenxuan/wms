package main

import (
	"flag"
	"wms/config"
	"wms/db"
	"wms/log"
	"wms/router"
)

func main() {
	var configFile string
	var err error

	flag.StringVar(&configFile, "config", "./config/cfg_dev.json", "config file")
	flag.Parse()

	if configFile == "" {
		log.Errorf("config file argument is required.")
		return
	}

	err = config.Init(configFile)
	if err != nil {
		log.Errorf("[init][config] %v", err.Error())
		return
	}

	err = db.InitMysql()
	if err != nil {
		log.Errorf("[init][mysql] %v", err.Error())
		return
	}

	err = db.InitRedis()
	if err != nil {
		log.Errorf("[init][redis] %v", err.Error())
	}

	router.Init()
}
