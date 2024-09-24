package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"petProject/internal/model"
	"petProject/internal/response"
	"strings"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if user.Name == "" {
		response.NewErrorResponse(w, http.StatusForbidden, errors.New("name is required"))
		return
	}
	if user.Password == "" {
		response.NewErrorResponse(w, http.StatusForbidden, errors.New("password is required"))
		return
	}

	responsePassword := strings.Repeat("*", len(user.Password))

	err := h.service.AuthorizationService.CheckIsUserExist(user.Name)
	if err != nil {
		response.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	user, err = h.service.AuthorizationService.CreateUser(user)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = responsePassword
	response.RespondJSON(w, user, http.StatusCreated)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input model.SignInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.Name == "" {
		response.NewErrorResponse(w, http.StatusForbidden, errors.New("name is required"))
		return
	}
	if input.Password == "" {
		response.NewErrorResponse(w, http.StatusForbidden, errors.New("password is required"))
		return
	}

	err := h.service.CashService.CheckTokenExist(input.Name, input.Device)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.AuthorizationService.GenerateToken(input.Name, input.Password)
	if err != nil {
		response.NewErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	err = h.service.AuthorizationService.WriteTokenInRedis(input.Name, token, input.Device)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondJSON(w, model.TokenResponse{
		token},
		http.StatusOK)
}
