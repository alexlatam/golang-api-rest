package handlers

import (
	"api-golang/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUsersByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUsersByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user = models.UpdateUser(user, userResponse)
	models.SendData(w, user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUsersByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.DeleteUser(user.Id)
		models.SendNoContent(w)
	}
}

func getUsersByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	if user, err := models.GetUser(userId); err != nil {
		return user, err
	} else {
		return user, nil
	}
}
