package user_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/user_usecase"
	"log"
	"net/http"
)

func PostUpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u model.UpdateUserDetails
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := user_usecase.GetUpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
