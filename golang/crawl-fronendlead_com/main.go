package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	URLProblems := "https://api.frontendlead.com/graphql"
	payloadStr := `{"query":"\n  {\n    problems(first: 250) {\n      nodes {\n        title\n\t\tcontent\n        slug\n        date\n        featuredImage {\n          node {\n            sourceUrl\n          }\n        }\n        info {\n            company\n            category\n            type\n            }\n        common {\n            difficulty\n            videoLink {\n                url\n            }\n            comingSoon\n            subtitle\n            premiumQuestion\n\t\t\tstudyPlan\n        }\n      }\n    }\n  }\n  "} `

	body := strings.NewReader(payloadStr)

	res, err := http.Post(URLProblems, "application/json", body)
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer res.Body.Close()

	// Check for HTTP status code
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", res.StatusCode)
	}

	// Read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// // Parse the JSON response
	// var jsonResponse map[string]interface{}
	// if err := json.Unmarshal(resBody, &jsonResponse); err != nil {
	// 	log.Fatalf("Error parsing JSON response: %v", err)
	// }
	// log.Printf("Response JSON:\n%v", jsonResponse)

	filename := "response.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 666)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	defer file.Close()

	file.Write(resBody)

	// // Pretty print the JSON response
	// var prettyJSON bytes.Buffer
	// if err := json.Indent(&prettyJSON, resBody, "", "  "); err != nil {
	// 	log.Fatalf("Error formatting JSON response: %v", err)
	// }

	// log.Printf("Response JSON:\n%s", prettyJSON.String())
}
