package post_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/post_usecase"
	"log"
	"net/http"
)

func PostUpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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
	if err := post_usecase.UpdatePost(p); err != nil {
		log.Printf("fail: post_usecase.UpdatePost, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
