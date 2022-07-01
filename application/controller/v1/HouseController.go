package v1

import (
	"backend-go-2/domain/usecases/house"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleGetHouse(writer http.ResponseWriter, request *http.Request) {
	var houseResponse = house.ExecuteGetHouseUseCase()

	jsonResponse, err := json.Marshal(&houseResponse)

	if err != nil {
		fmt.Println("Error occurred while creating JSON response: ", err)
		return // TODO: send error response
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(jsonResponse)
	if err != nil {
		fmt.Println("Error occurred while sending response: ", err)
	}
}

func AddHouseHandler(router *mux.Router) {
	router.HandleFunc("/v1/house", handleGetHouse).Methods("GET")
}
