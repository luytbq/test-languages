package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gihub.com/luytbq/frontendleaddotcom-mimic/util"
)

func handleRun(w http.ResponseWriter, r *http.Request) {
	log.Println("handle post submit")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req RunRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil || req.Submission == "" || req.Slug == "" {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	question, err := util.GetQuestionBySlug(req.Slug)
	if err != nil {
		log.Printf("Error reading question: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err = deletePreviousUserCode(); err != nil {
		log.Printf("Error cleaning previous test: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err = deletePreviousTestCases(); err != nil {
		log.Printf("Error cleaning previous test: %s", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = writeUserCode(req.Submission)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	for i := range question.TestCases {
		tc := question.TestCases[i]

		writeTestCase(tc.ID, tc.Test)
	}

	res := &RunResponse{State: "success"}
	w.Header().Add("Content-Type", "application/json")

	cmd := exec.Command("jasmine")
	cmd.Dir = "../test-engine/"

	log.Println("execute jasmine")
	outputBytes, err := cmd.CombinedOutput()

	if err != nil {
		res.State = "error"
	}

	log.Printf("error: %v", err)
	log.Printf("jasmine output: \n%s", string(outputBytes))

	res.Message = string(outputBytes)

	_ = json.NewEncoder(w).Encode(res)
}

func writeUserCode(submission string) error {
	filePath := "../test-engine/src/user-code.mjs"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Printf("Failed to open file: %s", err.Error())
		return err
	}
	_, err = file.WriteString(submission)
	if err != nil {
		log.Printf("Failed to write user code: %s", err.Error())
		return err
	}
	file.Close()
	return err
}

func writeTestCase(id int, testScript string) error {
	testScript = strings.ReplaceAll(testScript, ".to.deep.equal(", ".toEqual(")

	filePath := fmt.Sprintf("../test-engine/spec/test-cases/test-case-%d.js", id)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Printf("Failed to open file: %s", err.Error())
		return err
	}
	_, err = file.WriteString(testScript)
	if err != nil {
		log.Printf("Failed to write user code: %s", err.Error())
		return err
	}
	file.Close()
	return err
}

func deletePreviousUserCode() error {
	files, err := filepath.Glob("../test-engine/src/*.mjs")
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}

func deletePreviousTestCases() error {
	files, err := filepath.Glob("../test-engine/spec/test-cases/*.js")
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}
