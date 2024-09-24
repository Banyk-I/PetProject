package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"petProject/internal/model"
	"petProject/internal/response"
)

func (h *Handler) GetActiveAccessTokens(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*model.TokenClaims)
	token := r.Context().Value("token").(string)

	err := h.service.CashService.CheckUserAuthorized(token)
	if err != nil {
		response.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	err = h.service.AdminService.VerificationForAdmin(claims.UserID)
	if err != nil {
		response.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}
	logrus.Info("User ", claims.UserID, " is admin")

	userTokens, err := h.service.AdminService.GetActiveAccessTokens()
	if err != nil {
		response.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.RespondJSON(w, userTokens, http.StatusOK)
}

func (h *Handler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*model.TokenClaims)
	token := r.Context().Value("token").(string)

	err := h.service.CashService.CheckUserAuthorized(token)
	if err != nil {
		response.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	err = h.service.AdminService.VerificationForAdmin(claims.UserID)
	if err != nil {
		response.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	userLogout := r.URL.Query().Get("user")
	device := r.URL.Query().Get("device")

	if userLogout == "" {
		response.NewErrorResponse(w, http.StatusBadRequest, "user parameter is missing")
		return
	}

	err = h.service.AdminService.LogoutUser(userLogout, device)
	if err != nil {
		response.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	response.RespondJSON(w, response.MessageResponse{fmt.Sprintf("%s is logged out", userLogout)}, http.StatusOK)
}
