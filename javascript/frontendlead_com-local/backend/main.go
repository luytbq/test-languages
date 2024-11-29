package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"gihub.com/luytbq/frontendleaddotcom-mimic/handler"
)

func main() {
	// Serve the static HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle the /run endpoint
	http.HandleFunc("/run", handler.HandlePostSubmit)

	// Start the server
	fmt.Println("Server running at http://localhost:9080")
	log.Fatal(http.ListenAndServe(":9080", nil))
}

func execTest() {
	// Define the command to run
	cmd := exec.Command("node", "./test-engine/index.js")

	// Capture the standard output and standard error
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		fmt.Printf("Error details: %s\n", errBuf.String())
		return
	}

	// Print the output
	fmt.Println("Command Output:")
	fmt.Println(outBuf.String())
}
