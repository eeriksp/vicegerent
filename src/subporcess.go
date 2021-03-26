package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

// Execute a single command and log its stdout and stderr to the HTTP response.
func ExecuteCommand(response http.ResponseWriter, workDir string, commands ...string) {
	fmt.Fprintf(response, "--- Start to execute `%s` ---\n", commands)
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
