package common

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Auth struct {
	Token *jwtauth.JWTAuth `json:"token"`
}

func InitAuth(auth *Auth) {
	secret := jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(0))
	auth.Token = secret
	_, tokenString, _ := secret.Encode(map[string]interface{}{"user_id": 1})
	fmt.Println(tokenString)
}

func SessionAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println(cookie.Value)
		next.ServeHTTP(w, r)
	})
}
