package house

import (
	"backend-go-2/domain/model"
)

func ExecuteGetHouseUseCase() model.House {
	var house = model.House{
		Id:   "123",
		Name: "FlairPlatform",
	}

	return house
}
