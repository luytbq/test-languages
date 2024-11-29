package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"gihub.com/luytbq/frontendleaddotcom-mimic/types"
)

func HandlePostSubmit(w http.ResponseWriter, r *http.Request) {
	log.Println("handle post submit")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req types.RunRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil || req.Submission == "" {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	// Write user code to a temporary file in backend/tmp
	tempFileUserCode, err := os.CreateTemp(filepath.Join("../temp-file", "user-code"), "usercode-*.js")
	if err != nil {
		log.Printf("Failed to create temp file: %s", err.Error())
		http.Error(w, "Failed to create temp file", http.StatusInternalServerError)
		return
	}
	tempFileTestCase, err := os.CreateTemp(filepath.Join("../temp-file", "test-cases"), "test-case-*.js")
	if err != nil {
		log.Printf("Failed to create temp file: %s", err.Error())
		http.Error(w, "Failed to create temp file", http.StatusInternalServerError)
		return
	}

	defer tempFileUserCode.Close()
	// defer os.Remove(tempFileUserCode.Name())
	// defer os.Remove(tempFileTestCase.Name())

	_, err = tempFileUserCode.WriteString(req.Submission)
	if err != nil {
		http.Error(w, "Failed to write user code", http.StatusInternalServerError)
		return
	}

	_, err = tempFileTestCase.WriteString(`
		it('should handledeeply nested arrays', (done) => {
		const input = [[[[1]]], 2, [3, 4], 5];
		const expected = [1, 2, 3, 4];
		const result = flatternArray(input);
		expect(result).to.deep.equal(expected);
		done();
		})`)

	if err != nil {
		http.Error(w, "Failed to write user code", http.StatusInternalServerError)
		return
	}

	log.Printf("USER_CODE_PATH=%s", tempFileUserCode.Name())
	log.Printf("TEST_CASE_PATH=%s", tempFileTestCase.Name())

	// Run the test engine
	cmd := exec.Command("node", "index.js")
	cmd.Dir = "../test-engine" // Set working directory to test-engine
	cmd.Env = append(os.Environ(), fmt.Sprintf("USER_CODE_PATH=%s", tempFileUserCode.Name()))
	cmd.Env = append(os.Environ(), fmt.Sprintf("TEST_CASE_PATH=%s", tempFileTestCase.Name()))

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	log.Printf("Stdout: %s", outBuf.String())
	log.Printf("Stderr: %s", errBuf.String())

	err = cmd.Run()
	if err != nil {
		log.Printf("Error running test engine: %s", err.Error())
		resp := types.RunResponse{
			Output: "",
			Error:  errBuf.String(),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	// Return the test engine output
	resp := types.RunResponse{
		Output: outBuf.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
