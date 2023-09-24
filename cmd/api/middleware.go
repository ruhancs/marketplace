package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func (app *Application) Auth(next http.Handler) http.Handler {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			render.Status(r, 401)
			render.JSON(w,r, map[string]string{"error": "you must be logged in to access"})
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			render.Status(r, 401)
			render.JSON(w,r, map[string]string{"error": "no authorization header received"})
			return
		}

		token := headerParts[1]
		//marketplace-go nome do realm
		provider,err := oidc.NewProvider(r.Context(), os.Getenv("KEYCLOCK_URL"))
		if err != nil {
			render.Status(r,500)
			render.JSON(w,r, map[string]string{"error": "error to connect to identity provider"})
			return
		}
		
		verifier := provider.Verifier(&oidc.Config{ClientID: "marketplace"})
		//verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})
		_,err = verifier.Verify(r.Context(),token)
		if err != nil {
			render.Status(r,401)
			render.JSON(w,r, map[string]string{"error": "invalid token"})
			return
		}

		next.ServeHTTP(w,r)
	})
}