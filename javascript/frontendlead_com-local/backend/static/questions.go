package static

import (
	"encoding/json"
	"log"
	"os"

	"gihub.com/luytbq/frontendleaddotcom-mimic/types"
)

var Questions = loadQuestions()

func loadQuestions() []types.Question {
	filename := "./static/crawled-questions.json"
	log.Printf("Reading %s", filename)

	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("Open file error: %s", err.Error())
	}
	defer file.Close()

	res := make([]types.Question, 0)
	err = json.NewDecoder(file).Decode(&res)
	if err != nil {
		log.Fatalf("Decode questions error: %s", err.Error())
	}

	log.Printf("%d questions loaded", len(res))

	return res
}
