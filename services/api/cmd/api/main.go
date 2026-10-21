package main

import (
	"net/http"

	"github.com/Utkarsh0uchiha/pitchvision-ai/services/api/internal/health"
)

func main() {
	http.Handle("/healthz", health.HealthHandler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
