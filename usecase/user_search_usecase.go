package usecase

import (
	"hackathon/dao"
	"log"
	"net/http"
)

func GetSearchUser(w http.ResponseWriter, r *http.Request) (bytes []byte) {
	Name := r.URL.Query().Get("name")
	if Name == "" {
		log.Println("fail: name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bytes = dao.SearchUser(w, Name)
	return bytes
}
