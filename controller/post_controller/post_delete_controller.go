package post_controller

import (
	"hackathon/usecase/post_usecase"
	"log"
	"net/http"
)

func PostDeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	if err := post_usecase.DeletePost(id); err != nil {
		log.Printf("fail: post_usecase.DeletePost, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
