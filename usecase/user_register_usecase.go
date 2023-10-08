package usecase

import (
	"encoding/json"
	"github.com/oklog/ulid/v2"
	"hackathon/dao"
	"hackathon/model"
	"log"
	"net/http"
)

func GetRegisterUser(w http.ResponseWriter, r *http.Request) (bytes []byte) { // ①
	var u model.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := validateInput(w, u); err != "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf(err)
	}

	ulidId := ulid.Make()
	id := ulidId.String()

	dao.RegisterUser(w, u, id)

	var response model.Response
	response.Id = id

	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return bytes
}

func validateInput(w http.ResponseWriter, u model.User) (err string) {
	if u.Name == "" || len(u.Name) > 50 {
		err = "fail: invalid user name"
		return err
	}

	if u.Age < 20 || u.Age > 80 {
		err = "fail: invalid user age"
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	return ""
}
