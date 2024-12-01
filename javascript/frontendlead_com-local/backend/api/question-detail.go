package api

import (
	"encoding/json"
	"net/http"

	"gihub.com/luytbq/frontendleaddotcom-mimic/util"
)

func handleGetDetail(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if ok := query.Has("slug"); !ok {
		http.Error(w, "slug not provided", http.StatusBadRequest)
	}
	slug := query.Get("slug")

	questionDetail, err := util.GetQuestionBySlug(slug)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questionDetail)
}
