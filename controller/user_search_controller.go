package controller

import (
	"hackathon/usecase"
	"log"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		bytes := usecase.GetSearchUser(w, r)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(bytes)
		if err != nil {
			return
		}
	}
}
