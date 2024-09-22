package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

const minScore = 45
const maxScore = 100

type student_score struct {
  FinalScore int `json:"student_score"`
}

func getScore(w http.ResponseWriter, r *http.Request) {
  var finalScore = student_score{
      FinalScore: (rand.Intn(maxScore-minScore) + minScore),
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(finalScore)
}

func handleRequests() {
  http.Handle("/creditscore", http.HandlerFunc(getScore))
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
  handleRequests()
}