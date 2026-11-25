package server

import (
	"net/http"

	"github.com/Utkarsh0uchiha/pitchvision-ai/services/api/internal/health"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/healthz", health.HealthHandler())

	return mux
}
