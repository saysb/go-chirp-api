package handlers

import (
	"encoding/json"
	"net/http"
	"twitter-clone-api/internal/models"
	"twitter-clone-api/internal/services"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var userRequest models.CreateUserRequest


    if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
        Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    user := models.User{
        Username: userRequest.Username,
        Email:    userRequest.Email,
        Password: userRequest.Password,
    }

    if err := h.service.Create(r.Context(), &user); err != nil {
        switch err.Error() {
        case "userAlreadyExists":
            Error(w, http.StatusBadRequest, "userAlreadyExists")
        default:
            Error(w, http.StatusInternalServerError, err.Error())
        }
        return
    }

    Success(w, http.StatusCreated, user)
}

func (h *UserHandler) GetOne(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    user, err := h.service.GetByID(r.Context(), id)
    if err != nil {
        Error(w, http.StatusNotFound, "User not found")
        return
    }

    Success(w, http.StatusOK, user)
}


func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := h.service.Update(r.Context(), id, &user); err != nil {
        Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    Success(w, http.StatusOK, user)
}