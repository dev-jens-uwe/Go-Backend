package main

import (
	"backend-go-2/application"
	"fmt"
	"net/http"
)

func main() {
	router := application.GetRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Port already in use")
	}
}
