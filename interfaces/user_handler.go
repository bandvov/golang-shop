package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain/users"
	"github.com/go-playground/validator"
)

type UserHandler struct {
	UserService *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	users, err := h.UserService.GetUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r)
	id, err := strconv.Atoi(r.PathValue("id"))
	fmt.Println("err", err)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	user, err := h.UserService.GetUserByID(ctx, id)
	if err != nil {
		if err == users.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	// Validate the User struct
	err := validate.Struct(user)
	if err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.UserService.CreateUser(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.UserService.UpdateUser(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.UserService.DeleteUser(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
