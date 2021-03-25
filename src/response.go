package main

import (
	"fmt"
	"net/http"
)

func AppendCommandOutputToResponse(response http.ResponseWriter, command []string, output string, err error) {
	fmt.Fprintf(response, "--- Start to execute `%s` ---\n", command)
	fmt.Fprintln(response, output)
	if err != nil {
		fmt.Fprintf(response,"ERROR: %s", err)
	}
}
