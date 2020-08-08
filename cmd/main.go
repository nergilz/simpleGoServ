package main

import (
	"os"

	"io"
	"log"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// AppKey ...
const AppKey = "secret"

// main func...
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/status", StatusHandler).Methods("GET")
	router.HandleFunc("/token", TokenHandler).Methods("GET")
	router.Handle("/", AuthMiddleware(http.HandlerFunc(ExampleHandler)))

	http.ListenAndServe(":8001", handlers.LoggingHandler(os.Stdout, router))

}

// StatusHandler ...
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("my Not Implemented /status"))
}

// ExampleHandler ...
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

// TokenHandler ...
func TokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application-json")
	// r.ParseForm()
	username := "Tom Cruse" // r.Form.Get("username")
	//password := "qwerty"    // r.Form.Get("password")
	// if username != "myusername" || password != "mypassword" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	io.WriteString(w, `{"error":"invalid_credentials"}`)
	// 	return
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"role": "admin",
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(AppKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error!":"token_generation_faled"}`)
		return
	}

	//io.WriteString(w, `{"token":"`+tokenString+`"}`)
	w.Write([]byte(tokenString))
	return
}

// AuthMiddleware ...
func AuthMiddleware(next http.Handler) http.Handler {
	if len(AppKey) == 0 {
		log.Fatal("HTTP server unable to start, expected an AppKey for JWT auth")
	}

	jwtmiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(AppKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtmiddleware.Handler(next)
}
