package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/database"
	"github.com/fernandomocrosky/DevBookGo/src/models"
	"github.com/fernandomocrosky/DevBookGo/src/repositories"
	"github.com/fernandomocrosky/DevBookGo/src/responses"
	"github.com/fernandomocrosky/DevBookGo/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	storedUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.CheckHash(storedUser.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	responses.JSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Logged in sucessfully",
	})
}
