package controller

import (
	"hackathon/usecase"
	"log"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes := usecase.GetRegisterUser(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(bytes)
	if err != nil {
		return
	}
}
