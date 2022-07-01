package v1

import (
	"backend-go-2/domain/usecases/house"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleGetHouse(writer http.ResponseWriter, request *http.Request) {
	var houseResponse = house.ExecuteGetHouseUseCase()

	err := SendJSONResponse(writer, houseResponse)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func AddHouseHandler(router *mux.Router) {
	router.HandleFunc("/v1/house", handleGetHouse).Methods("GET")
}
