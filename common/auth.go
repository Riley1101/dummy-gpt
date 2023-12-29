package common

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const (
	Algorithms = "HS256"
	Secret     = "secret"
)

type Auth struct {
	Token *jwtauth.JWTAuth `json:"token"`
}

func InitAuth(auth *Auth) {
	secret := jwtauth.New(Algorithms, []byte(Secret), nil, jwt.WithAcceptableSkew(0))
	auth.Token = secret
	_, tokenString, _ := secret.Encode(map[string]interface{}{"user_id": 1})
	log.Println(tokenString)
}

func SessionAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			// TODO: design choice will update this better in the future but now just redirect
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		log.Println(cookie.Value)
		next.ServeHTTP(w, r)
	})
}
