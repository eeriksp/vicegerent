package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Address string
	Tasks   []Task
}

type Task struct {
	Url      string
	WorkDir  string
	Commands [][]string
}

func (task Task) fullUrl() string {
	return urlPrefix + task.Url
}

func loadConfiguration() Config {
	config := Config{}
	err := json.Unmarshal(loadConfigurationFromFile(), &config)
	panicIf(err)
	return config
}

func loadConfigurationFromFile() []byte {
	content, err := ioutil.ReadFile(configFilePath)
	panicIf(err)
	return content
}
