package user_controller

import (
	"hackathon/usecase/user_usecase"
	"log"
	"net/http"
)

func PostDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := user_usecase.GetDeleteUser(id); err != nil {
		log.Printf("fail: GetDeleteUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
