package v1

import (
	"api-golang/models"
	"encoding/json"
	"errors"
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

	// Verificamos que el body del request sea valido
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	// Verificamos que los datos del usuario a crear sean validos
	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.SetPassword(user.Password)
	// Verficamos que el usuario se haya creado correctamente
	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	models.SendData(w, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUsersByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	request := &models.User{}
	decoder := json.NewDecoder(r.Body)

	// Verificamos que el body del request sea valido
	if err := decoder.Decode(request); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	// Verificamos que los datos del usuario a crear sean validos
	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.SetPassword(user.Password)

	// Verficamos que el usuario se haya actualizado correctamente
	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	models.SendData(w, user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUsersByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		user.Delete()
		models.SendNoContent(w)
	}
}

func getUsersByRequest(r *http.Request) (*models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.GetUserById(userId)

	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}
