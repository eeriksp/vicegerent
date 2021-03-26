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
		fmt.Fprintf(response, `Started task "%s"`, task.Name)
		fmt.Fprintln(response, "\n")
		fmt.Fprintln(response)
		for _, command := range task.Commands {
			ExecuteCommand(response, task.WorkDir, command...)
			tryFlush(response)
		}
		fmt.Fprintln(response, "\nTask finished.")

	}
}

func registerHandlers(config Config) {
	for _, task := range config.Tasks {
		http.HandleFunc(task.fullUrl(), createHandler(task))
	}
}

func main() {
	config := loadConfiguration()
	fmt.Println("Ready to accept tasks. Listening for")
	for _, task := range config.Tasks {
		fmt.Println("- " + task.Name + ": " + task.fullUrl())
	}
	registerHandlers(config)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}
