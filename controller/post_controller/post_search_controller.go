package post_controller

import (
	"hackathon/usecase/post_usecase"
	"log"
	"net/http"
	"strconv"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	bytes, err := post_usecase.GetPost(id)

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
	stringId := r.URL.Query().Get("id")
	id, convertErr := strconv.Atoi(stringId)
	if convertErr != nil {
		log.Printf("fail: strconv.Atoi, %v\n", convertErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := post_usecase.GetAllPostsByCategory(id)

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
	id := r.URL.Query().Get("id")
	bytes, err := post_usecase.GetAllPostsByUser(id)

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
	stringId := r.URL.Query().Get("id")
	id, convertErr := strconv.Atoi(stringId)
	if convertErr != nil {
		log.Printf("fail: strconv.Atoi, %v\n", convertErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := post_usecase.GetAllPostsByCurriculum(id)

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
	id := r.URL.Query().Get("id")
	bytes, err := post_usecase.GetAllPostsByTag(id)

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
