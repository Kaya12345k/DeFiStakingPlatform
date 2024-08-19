package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Stake struct {
	ID     int     `json:"id"`
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
}

var stakes = []Stake{
	{ID: 1, UserID: "123", Amount: 100.50},
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	http.HandleFunc("/createStake", createStakeHandler)
	http.HandleFunc("/getStakes", getStakesHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}

func createStakeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newStake Stake
	err := json.NewDecoder(r.Body).Decode(&newStake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newStake.ID = len(stakes) + 1
	stakes = append(stakes, newStake)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newStake)
}

func getStakesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stakes)
}