package user_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/user_usecase"
	"log"
	"net/http"
)

func PostDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
