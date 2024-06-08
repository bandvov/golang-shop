package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain"
)

type UserHandler struct {
    UserService *application.UserService
}
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.GetUsers(w, r)
    case http.MethodPost:
        h.CreateUser(w, r)
    case http.MethodPut:
        h.UpdateUser(w, r)
    case http.MethodDelete:
        h.DeleteUser(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.UserService.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := h.UserService.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user domain.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    if err := h.UserService.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    var user domain.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    if err := h.UserService.UpdateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    if err := h.UserService.DeleteUser(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}