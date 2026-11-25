package main

import (
	"net/http"
	"github.com/Utkarsh0uchiha/pitchvision-ai/services/api/internal/server"
)

func main() {

	mux := server.NewRouter()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
