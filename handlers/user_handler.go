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
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.userRepo.CreateUser(&newUser, newUser.PasswordDigest)
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
