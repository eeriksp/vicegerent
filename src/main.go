package main

import (
	"fmt"
	"log"
	"net/http"
)

func tryFlush(response http.ResponseWriter) {
	if f, ok := response.(http.Flusher); ok {
		f.Flush()
	}
}

func createHandler(task Task) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		for _, command := range task.Commands {
			ExecuteCommand(response, task.WorkDir, command...)
			tryFlush(response)
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
