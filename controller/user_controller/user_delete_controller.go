package user_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/user_usecase"
	"log"
	"net/http"
)

func PostDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decode := json.NewDecoder(r.Body)
	var i model.SearchById
	err := decode.Decode(&i)
	if err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := user_usecase.GetDeleteUser(i); err != nil {
		log.Printf("fail: GetDeleteUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
