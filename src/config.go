package main

import (
	"gopkg.in/yaml.v3"
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
	err := yaml.Unmarshal(loadConfigurationFromFile(), &config)
	panicIf(err)
	return config
}

func loadConfigurationFromFile() []byte {
	content, err := ioutil.ReadFile(configFilePath)
	panicIf(err)
	return content
}
