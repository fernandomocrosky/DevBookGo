package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/auth"
	"github.com/fernandomocrosky/DevBookGo/src/database"
	"github.com/fernandomocrosky/DevBookGo/src/models"
	"github.com/fernandomocrosky/DevBookGo/src/repositories"
	"github.com/fernandomocrosky/DevBookGo/src/responses"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Posts
	if err := json.Unmarshal(body, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	post.UserId = userId

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewPostsRepository(db)
	post.ID, err = repo.CreatePost(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {}

func GetPostbyID(w http.ResponseWriter, r *http.Request) {}

func UpdatePost(w http.ResponseWriter, r *http.Request) {}

func DeletePost(w http.ResponseWriter, r *http.Request) {}
