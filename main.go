package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type SpinResult struct {
	Result    []string `json:"result"`
	WinType   string   `json:"winType"`
	WinAmount int      `json:"winAmount"`
}

var symbols = []string{"🍒", "🍋", "🍊", "🍑", "🔔", "🍫", "7️⃣"}

func getRandomSymbols() []string {
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 7)
	for i := 0; i < 7; i++ {
		result[i] = symbols[rand.Intn(len(symbols))]
	}
	return result
}

func calculateWin(result []string) (string, int) {
	counts := make(map[string]int)
	for _, symbol := range result {
		counts[symbol]++
	}

	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	switch maxCount {
	case 7:
		return "Jackpot", 1000
	case 6:
		return "Six of a kind", 500
	case 5:
		return "Five of a kind", 100
	case 4:
		return "Four of a kind", 50
	case 3:
		return "Three of a kind", 20
	case 2:
		return "Two of a kind", 10
	default:
		return "No win", 0
	}
}

func spinHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	result := getRandomSymbols()
	winType, winAmount := calculateWin(result)

	spinResult := SpinResult{
		Result:    result,
		WinType:   winType,
		WinAmount: winAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spinResult)
}

func main() {
	http.HandleFunc("/spin", spinHandler)
	http.ListenAndServe(":8080", nil)
}
