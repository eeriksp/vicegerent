package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

// createHandler returns a function which can be registered with `http.HandleFunc`
// and which takes care of executing all commands of the given `task`
// and reporting the output to the HTTP response.
func createHandler(task Task) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, `Started task "%s"`, task.Name)
		fmt.Fprintln(response, "\n\n")
		for _, command := range task.Commands {
			executeCommand(response, task.WorkDir, command...)
		}
		fmt.Fprintln(response, "\nTask finished.")

	}
}

// executeCommand executes a single command and logs its stdout and stderr to the HTTP response.
func executeCommand(response http.ResponseWriter, workDir string, commands ...string) {
	fmt.Fprintf(response, "--- Start to execute %s ---\n", commands)
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Dir = workDir
	cmd.Stdout = response
	cmd.Stderr = response
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(response, "ERROR: %s", err)
	}
	fmt.Fprintln(response)
}
