package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Config represents the entire application configuration.
type Config struct {
	Address string // the address the application will be listening to, e.g. `:8080`
	Tasks   []Task // a list of tasks the application will be capable of performing.
}

// A Task represents one standalone unit of work that can be invoked by an HTTP request.
type Task struct {
	Name     string // a human readable name of the task
	Url      string // urlPrefix + Url as constructed by `task.fullUrl()` will be the URL the task can be invoked by
	WorkDir  string // the working directory for all the commands
	Commands [][]string
}

func (task Task) fullUrl() string {
	return urlPrefix + task.Url
}

// loadConfiguration loads the configuration from the disk and parses it into the Config struct.
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
