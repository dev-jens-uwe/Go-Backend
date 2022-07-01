package house

import "backend-go-2/domain/model"

func ExecuteFindHousesUseCase() []model.House {
	return []model.House{
		model.House{Id: "1", Name: "Flair"},
		model.House{Id: "1", Name: "Stadt-Villa"},
	}
}
