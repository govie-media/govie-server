package core

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenSalt = []byte("my_secret_key")

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

func ParseAuthRequest(r *http.Request) (*string, time.Time, error) {
	var user AuthUser

	// TODO: Handle Authorization Header
	// username, password, ok := r.BasicAuth()
	// Parse Form Request
	err := r.ParseForm()
	if err != nil {
		return nil, time.Now(), err
	}

	// Get the username and password from the form data
	username := r.FormValue("username")
	password := r.FormValue("password")
	remember := r.FormValue("rememberMe")

	user = AuthUser{
		username,
		password,
	}

	// Parse JSON Request
	/*
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			println("ERROR with USER")
			return nil, err
		}
	*/

	// Dummy credential check
	if validateUser(user) {
		token, expirationTime, err := generateToken(user.Username, 0, remember == "true")
		if err != nil {
			println("ERROR with TOKEN")
			return nil, time.Now(), err
		}

		return &token, expirationTime, nil
	}

	return nil, time.Now(), errors.New("invalid credentials")
}

func ValidateAuth(next http.HandlerFunc, redirect string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			// Check if the request has an Authorization header
			if r.Header.Get("Authorization") == "" {
				http.Redirect(w, r, redirect, http.StatusFound)
				return
			}
		*/

		/*
			// Extract the token from the Authorization header
			token := r.Header.Get("Authorization")
			// Validate the token using your favorite token validation library
			if !parseToken(token) {
				http.Redirect(w, r, redirect, http.StatusFound)
				return
			}
		*/

		// If the token is valid, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func ValidateCookieAuth(next http.HandlerFunc, redirect string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()

		if len(cookies) > 0 {
			for _, cookie := range cookies {
				if cookie.Name == "gid" {
					token, err := parseToken(cookie.Value)

					if err == nil && token != nil {
						next.ServeHTTP(w, r)
						return
					}
				}
			}
		}

		http.Redirect(w, r, redirect, http.StatusFound)
		return
	})
}

func parseToken(tokenString string) (*jwt.Token, error) {
	// Create a new JWT token
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check if the token is signed with the expected method
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("Invalid Token")
		}

		// Return the secret key
		return tokenSalt, nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err)
		return nil, errors.New("Invalid Token")
	}

	return token, nil
}

func generateToken(username string, userId int, isLongLivedToken bool) (string, time.Time, error) {
	expirationTime := getTokenExpirationTime(isLongLivedToken)

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthClaims{
		10,
		jwt.RegisteredClaims{
			Issuer:    "govie",
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        strconv.Itoa(userId),
		},
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString(tokenSalt)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func comparePasswordHash(storedHash, newHash []byte) bool {
	// Compare the two hashes
	for i := range storedHash {
		if storedHash[i] != newHash[i] {
			return false
		}
	}
	return true
}

// TODO: pull user from db
func validateUser(user AuthUser) bool {
	return true
}

func getTokenExpirationTime(isLongLivedToken bool) time.Time {
	if isLongLivedToken {
		return time.Now().Add(30 * 24 * time.Hour)
	} else {
		return time.Now().Add(2 * time.Hour)
	}
}
