package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(response http.ResponseWriter, r *http.Request) {
	commands := loadConfiguration().Tasks[0].Commands
	for _, command := range commands {
		output, err := ExecuteCommand(command...)
		AppendCommandOutputToResponse(response, command, output, err)
	}
}

func createHandler(task Task) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		for _, command := range task.Commands {
			output, err := ExecuteCommand(command...)
			AppendCommandOutputToResponse(response, command, output, err)
		}
	}
}

func registerHandlers(config Config) {
	for _, task := range config.Tasks {
		http.HandleFunc(task.fullUrl(), createHandler(task))
	}
}

func main() {
	config := loadConfiguration()
	fmt.Println("Ready to accept tasks")
	fmt.Printf("%s", config)
	registerHandlers(config)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}
