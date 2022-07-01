package application

import (
	"backend-go-2/application/controller/v1"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	var router = mux.NewRouter()

	v1.AddAuthenticationHandler(router)
	v1.AddHouseHandler(router)

	return router
}
