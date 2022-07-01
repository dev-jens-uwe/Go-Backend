package v1

import (
	"backend-go-2/application/domain/usecases/house"
	"github.com/gorilla/mux"
	"net/http"
)

func handleGetHouse(writer http.ResponseWriter, request *http.Request) {
	house.ExecuteGetHouseUseCase()

	// TODO: return house via json
}

func AddHouseHandler(router *mux.Router) {
	router.HandleFunc("/v1/house", handleGetHouse).Methods("GET")
}
