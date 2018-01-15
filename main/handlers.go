package main

import(
		"encoding/json"
		"net/http"
		"strconv"
)

func submitScore(w http.ResponseWriter, r *http.Request){
	var submittedScore Score

	naam := r.FormValue("naam")
	score := r.FormValue("score")

	i, _ := strconv.ParseInt(score, 10, 64)

	submittedScore.Name = naam
	submittedScore.Score = i

	submitScoreToDatabase(submittedScore)
}

func showScore(w http.ResponseWriter, r *http.Request){
	scores := getScore()

	js, err := json.Marshal(scores)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}