package v1

import (
	"backend-go-2/application/response"
	usecase "backend-go-2/domain/usecases/house"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleGetHouse(writer http.ResponseWriter, request *http.Request) {
	var houseModel = usecase.ExecuteGetHouseUseCase()
	var getHouseResponse = response.SingleResourceResponse{Resource: &houseModel}

	err := SendJSONResponse(writer, getHouseResponse)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func handleFindHousesResponse(writer http.ResponseWriter, request *http.Request) {
	var houseModels = usecase.ExecuteFindHousesUseCase()
	var findHousesResponse = response.BulkResourceResponse{Resources: houseModels}

	err := SendJSONResponse(writer, findHousesResponse)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func AddHouseHandler(router *mux.Router) {
	router.HandleFunc("/v1/house/{id}", handleGetHouse).Methods("GET")
	router.HandleFunc("/v1/house", handleFindHousesResponse).Methods("GET")
}
