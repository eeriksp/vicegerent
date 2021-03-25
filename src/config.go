package main

import (
	"encoding/json"
	"io/ioutil"
)

const CofigFilename = "./config.json"

type Config struct {
	Port     int
	Commands [][]string
}

func loadConfiguration() Config {
	config := Config{}
	err := json.Unmarshal(loadConfigurationFromFile(), &config)
	if err != nil {
		panic(err)
	}
	return config
}

func loadConfigurationFromFile() []byte {
	content, err := ioutil.ReadFile(CofigFilename)
	if err != nil {
		panic(err)
	}
	return content
}
