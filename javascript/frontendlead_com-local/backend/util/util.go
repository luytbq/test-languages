package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gihub.com/luytbq/frontendleaddotcom-mimic/types"
)

func GetQuestionBySlug(slug string) (*types.QuestionDetail, error) {
	filename := fmt.Sprintf("./static/questions/%s.json", slug)
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("error opening file: %s", err.Error())
		return nil, err
	}

	questionDetail := &types.QuestionDetail{}

	err = json.NewDecoder(file).Decode(questionDetail)
	if err != nil {
		log.Printf("error decoding file: %s", err.Error())
		return nil, err
	}

	return questionDetail, nil
}
