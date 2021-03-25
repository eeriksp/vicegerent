package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const URL = "foo-bar"

func handler(response http.ResponseWriter, r *http.Request) {
	commands := loadConfiguration().Commands
	for _, command := range commands {
		output, err := ExecuteCommand(command...)
		AppendCommandOutputToResponse(response, command, output, err)
	}
}

func main() {
	config := loadConfiguration()
	fmt.Println("Here we are!")
	http.HandleFunc("/", handler)
	address := ":" + strconv.Itoa(config.Port)
	log.Fatal(http.ListenAndServe(address, nil))
}
