package user_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/user_usecase"
	"log"
	"net/http"
)

func PostNewUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u model.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := user_usecase.GetRegisterUser(u); err != nil {
		log.Printf("fail: GetRegisterUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
