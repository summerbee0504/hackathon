package post_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/post_usecase"
	"log"
	"net/http"
)

func PostUpdatePostHandler(w http.ResponseWriter, r *http.Request) {
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

	decoder := json.NewDecoder(r.Body)
	var p model.Post
	if err := decoder.Decode(&p); err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := post_usecase.UpdatePost(p)
	if err != nil {
		log.Printf("fail: post_usecase.UpdatePost, %v\n", err)
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
