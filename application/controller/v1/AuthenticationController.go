package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleLogin(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)

	_, err := responseWriter.Write([]byte("You are now logged in"))

	if err != nil {
		fmt.Println(err)
	}
}

func AddAuthenticationHandler(router *mux.Router) {
	router.HandleFunc("v1/authentication/login", handleLogin).Methods("POST")
}
