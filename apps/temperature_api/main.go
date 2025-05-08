package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// TemperatureResponse представляет ответ с температурой.
type TemperatureResponse struct {
	Location    string  `json:"location"`
	Temperature float64 `json:"temperature"`
}

// getRandomTemperature генерирует случайное значение температуры.
func getRandomTemperature() float64 {
	rand.Seed(time.Now().UnixNano())

	return rand.Float64()*100 - 50 // Температура в диапазоне от -50 до 50
}

// handleTemperature обрабатывает запросы на получение температуры.
func handleTemperature(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		http.Error(w, "Location is required", http.StatusBadRequest)

		return
	}

	temperature := getRandomTemperature()
	response := TemperatureResponse{
		Location:    location,
		Temperature: temperature,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"location": "%s", "temperature": %.2f}`, response.Location, response.Temperature)
}

// main запускает HTTP-сервер.
func main() {
	http.HandleFunc("/temperature", handleTemperature)
	fmt.Println("Сервер запущен на порту 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v\n", err)
	}
}
