package core

import (
	"net/http"
)

func HttpAuth(next http.HandlerFunc, redirect string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has an Authorization header
		if r.Header.Get("Authorization") == "" {
			http.Redirect(w, r, redirect, http.StatusFound)
			return
		}

		// Extract the token from the Authorization header
		token := r.Header.Get("Authorization")
		// Validate the token using your favorite token validation library
		if !validateToken(token) {
			http.Redirect(w, r, redirect, http.StatusFound)
			return
		}

		// If the token is valid, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func validateToken(token string) bool {
	/*
		// Parse the token
		tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key_here"), nil
		})
		if err != nil {
			return false
		}

		// Check if the token is valid
		if !tokenClaims.Valid {
			return false
		}

		// Get the claims from the token
		claims, ok := tokenClaims.(*jwt.StandardClaim)
		if !ok {
			return false
		}

		// Check if the token has expired
		if claims.ExpiresAt < time.Now().Unix() {
			return false
		}

		// Check if the token is for the correct user
		if claims.Subject != "your_user_id_here" {
			return false
		}

	*/

	return true
}
