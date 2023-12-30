package handlers

import (
	"encoding/json"
	"net/http"

	"pnlyzer/models"
)

type UserHandler struct {
	userRepo models.UserRepositoryGorm
}

func NewUserHandler(userRepo models.UserRepositoryGorm) *UserHandler {
	return &UserHandler{userRepo}
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	//var newUser models.User
	var userForm models.UserForm

	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newUser := models.User{
		Username: userForm.Username,
		Email:    userForm.Email,
	}

	err = h.userRepo.CreateUser(&newUser, userForm.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var userCreds models.UserCreds

	err := json.NewDecoder(r.Body).Decode(&userCreds)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.SignIn(userCreds.Username, userCreds.Password)

	if err != nil {
		http.Error(w, "Failed to authenticate user", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, "Failed to authenticate user", http.StatusUnauthorized)
		return
	}
}
