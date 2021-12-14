package usecases

import (
	"cleango/app/domains"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type UserInteractor struct {
	UserRepository UserRepository
	JSONResponse   JSONResponse
}

func (up *UserInteractor) Store(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := map[string]interface{}{
			"err": "body error",
		}
		res, _ := json.Marshal(msg)
		up.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, res)
		return
	}
	user := domains.User{
		ID: randString(10),
	}
	if err = json.Unmarshal(body, &user); err != nil {
		msg := map[string]interface{}{
			"err": "unmarshal error cause: " + err.Error(),
		}
		res, _ := json.Marshal(msg)
		up.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, res)
		return
	}
	log.Println(&user)
	_, err = up.UserRepository.Save(user)
	if err != nil {
		msg := map[string]interface{}{
			"err": "database error cause: " + err.Error(),
		}
		res, _ := json.Marshal(msg)
		up.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, res)
		return
	}
	msg, err := json.Marshal(user)
	if err != nil {
		up.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	up.JSONResponse.HTTPStatus(w, 201, msg)
}

func (up *UserInteractor) List(w http.ResponseWriter, r *http.Request) {
	users, err := up.UserRepository.FindAll()
	if err != nil {
		msg := map[string]interface{}{
			"err": "data error cause: " + err.Error(),
		}
		res, _ := json.Marshal(msg)
		up.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, res)
		return
	}
	msg, err := json.Marshal(users)
	if err != nil {
		up.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	up.JSONResponse.HTTPStatus(w, 200, msg)
}

func randString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}
