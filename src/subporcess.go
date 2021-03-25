package main

import "os/exec"

func ExecuteCommand(args ...string) (string, error) {
	command := exec.Command(args[0], args[1:]...)
	output, err := command.Output()
	return string(output), err
}
