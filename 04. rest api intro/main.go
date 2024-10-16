package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const minScore = 25
const maxScore = 100

type student_score struct {
	FinalScore int    `json:"student_score"`
	Grade      string `json:"grade"`
}

func getGrade(score uint) string {
	x := ""

	if score >= 70 && score <= 100 {
		x = "A"
	} else if score >= 60 && score < 70 {
		x = "B"
	} else if score >= 50 && score < 60 {
		x = "C"
	} else if score >= 45 && score < 50 {
		x = "D"
	} else if score >= 40 && score < 45 {
		x = "E"
	} else if score < 40 {
		x = "F"
	}

	return x
}

func getScore(w http.ResponseWriter, r *http.Request) {
	var score = rand.Intn(maxScore-minScore) + minScore
	var finalScore = student_score{
		FinalScore: score,
		Grade:      getGrade(uint(score)),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalScore)
}

// END POINT
func handleRequests() {
	http.Handle("/score", http.HandlerFunc(getScore))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
