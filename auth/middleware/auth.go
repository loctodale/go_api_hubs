package middleware

import (
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/loctodale/go_api_hubs_microservice/auth/global"
	"net/http"
	"strings"
)

type ContextKey string

const UserContextKey ContextKey = "user"

var (
	jwksSet jwk.Set
)

func InitJWKSFetcher(jwkUri string) error {
	set, err := jwk.Fetch(global.Ctx, jwkUri)
	if err != nil {
		return err
	}

	jwksSet = set

	fmt.Println("JWKS fetched successfully.")
	return nil
}

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
				return
			}
			tokenString := parts[1]

			token, err := jwt.Parse([]byte(tokenString), jwt.WithKeySet(jwksSet), jwt.WithValidate(true))
			if err != nil {
				http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
				return
			}

			_, ok := token.Get("user_id")
			if !ok {
				http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			}
		})
	}
}
