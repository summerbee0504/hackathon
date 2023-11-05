package post_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/post_usecase"
	"log"
	"net/http"
)

func PostLikePostHandler(w http.ResponseWriter, r *http.Request) {
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
	var l model.Like
	err := decode.Decode(&l)
	if err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := post_usecase.LikePost(l); err != nil {
		log.Printf("fail: post_usecase.LikePost, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func PostUnlikePostHandler(w http.ResponseWriter, r *http.Request) {
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
	var l model.Unlike
	err := decode.Decode(&l)
	if err != nil {
		log.Printf("fail: json.Decode, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errMsg := post_usecase.UnlikePost(l)
	if errMsg != nil {
		log.Printf("fail: post_usecase.UnlikePost, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetLikedPostByUserHandler(w http.ResponseWriter, r *http.Request) {
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

	posts, err := post_usecase.GetLikedPostByUser(i)
	if err != nil {
		log.Printf("fail: post_usecase.GetLikedPostByUser, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(posts); err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetLikeCountHandler(w http.ResponseWriter, r *http.Request) {
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

	count, err := post_usecase.GetLikeCount(i)
	if err != nil {
		log.Printf("fail: post_usecase.GetLikeCount, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(count); err != nil {
		log.Printf("fail: w.Write, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
