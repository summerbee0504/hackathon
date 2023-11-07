package user_controller

import (
	"hackathon/model"
	"hackathon/usecase/user_usecase"
	"log"
	"net/http"
)

func GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodGet {
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := user_usecase.ShowUserList()
	if err != nil {
		log.Printf("fail: user_usecase.ShowUserList, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(bytes); err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetUserDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var i model.SearchById
	i.Id = r.URL.Query().Get("id")
	bytes, err := user_usecase.GetShowUserDetail(i)

	if err != nil {
		log.Printf("fail: user_usecase.GetShowUserDetail, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(bytes); err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
