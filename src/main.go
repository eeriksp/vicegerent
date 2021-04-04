package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := loadConfiguration()
	printTasks(config)
	registerHandlers(config)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}

func printTasks(config Config) {
	fmt.Println("Ready to accept tasks. Listening on " + config.Address + " for")
	for _, task := range config.Tasks {
		fmt.Println("- " + task.Name + ": " + task.fullUrl())
	}
}

func registerHandlers(config Config) {
	for _, task := range config.Tasks {
		http.HandleFunc(task.fullUrl(), createHandler(task))
	}
}
