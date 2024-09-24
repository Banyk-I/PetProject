package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"petProject/internal/middleware"
	"petProject/internal/model"
	"petProject/internal/response"
	"petProject/internal/service"
)

type Handler struct {
	router  chi.Router
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	h := &Handler{
		router:  chi.NewRouter(),
		service: service,
	}
	h.InitRoutes()
	return h
}

func (h *Handler) InitRoutes() {
	h.router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})
	h.router.Route("/api", func(r chi.Router) {
		r.Use(middleware.AuthorizationMiddleware)

		r.Get("/test1", func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value("token").(string)

			err := h.service.CashService.CheckUserAuthorized(token)
			if err != nil {
				response.NewErrorResponse(w, http.StatusForbidden, err.Error())
				return
			}
			w.Write([]byte("Hello World, your token and claims"))
			fmt.Println(token)
			claims := r.Context().Value("claims").(*model.TokenClaims)
			fmt.Println(claims)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/access-tokens", h.GetActiveAccessTokens)
			r.Delete("/logout", h.LogoutUser)
		})
	})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
