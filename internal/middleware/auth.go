package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"petProject/internal/model"
	"petProject/internal/response"
	"strings"
)

func AuthorizationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			response.NewErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			response.NewErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if headerParts[0] != "Bearer" {
			response.NewErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if headerParts[1] == "" {
			response.NewErrorResponse(w, http.StatusUnauthorized, "token is empty")
			return
		}

		claims, err := parseToken(headerParts[1])
		if err != nil {
			response.NewErrorResponse(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		ctx = context.WithValue(ctx, "token", headerParts[1])

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func parseToken(accessToken string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(viper.GetString("secretKey")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}
	return claims, nil
}
