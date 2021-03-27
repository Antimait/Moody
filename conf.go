package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	dataFolder = "data"
	configFile = "conf.json"
)

type Conf struct {
	ConfPath     string                 `json:"conf_path" validate:"nonzero"`
	DbName       string                 `json:"db_name" validate:"nonzero"`
	ServerPort   string                 `json:"server_port" validate:"nonzero"`
	ProtocolConf map[string]interface{} `json:"protocols"`
}

func mustInitConf() Conf {
	var conf Conf
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fullPath := filepath.Join(pwd, dataFolder, configFile)
	file, fileErr := ioutil.ReadFile(fullPath)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	if jsonErr := json.Unmarshal(file, &conf); jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return conf
}
