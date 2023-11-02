package post_controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase/post_usecase"
	"log"
	"net/http"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var i model.SearchById
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := post_usecase.GetPost(i)

	if err != nil {
		log.Printf("fail: post_usecase.GetPost, %v\n", err)
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

func GetAllCurriculumsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes, err := post_usecase.GetAllCurriculums()

	if err != nil {
		log.Printf("fail: post_usecase.GetAllCurriculums, %v\n", err)
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

func GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes, err := post_usecase.GetAllTags()

	if err != nil {
		log.Printf("fail: post_usecase.GetAllTags, %v\n", err)
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

func GetPostsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var i model.SearchById
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes, err := post_usecase.GetAllPostsByCategory(i)

	if err != nil {
		log.Printf("fail: post_usecase.GetAllPostsByCategory, %v\n", err)
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

func GetPostsByUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var i model.SearchById
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	bytes, err := post_usecase.GetAllPostsByUser(i)

	if err != nil {
		log.Printf("fail: post_usecase.GetAllPostsByUser, %v\n", err)
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

func GetPostsByCurriculumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var i model.SearchById
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	bytes, err := post_usecase.GetAllPostsByCurriculum(i)

	if err != nil {
		log.Printf("fail: post_usecase.GetAllPostsByCurriculum, %v\n", err)
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

func GetPostsByDateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes, err := post_usecase.GetAllPostsByDate()

	if err != nil {
		log.Printf("fail: post_usecase.GetAllPostsByDate, %v\n", err)
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

func GetPostsByTagHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var i model.SearchById
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	bytes, err := post_usecase.GetAllPostsByTag(i)

	if err != nil {
		log.Printf("fail: post_usecase.GetAllPostsByTag, %v\n", err)
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
