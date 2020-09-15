package queries

import (
	"github.com/thiagomotadev/tmdgolangbase"
	"github.com/thiagomotadev/tmdgolangexample/dtos"
	"github.com/thiagomotadev/tmdgolangexample/entities"
	"github.com/thiagomotadev/tmdgolangexample/repositories"
)

func randomTextToDTO(entity entities.RandomText) (dto dtos.RandomText) {
	dto = dtos.RandomText{
		ID:      entity.ID,
		Text:    entity.Text,
		IsValid: entity.TextSum == tmdgolangbase.SumSHA256(entity.Text),
	}
	return
}

// RandomTextGetAll ...
func RandomTextGetAll(repository *repositories.RandomText) (allDTOs []dtos.RandomText, err error) {
	allEntities, err := repository.GetAll()
	allDTOs = make([]dtos.RandomText, 0)
	for _, item := range allEntities {
		allDTOs = append(allDTOs, randomTextToDTO(item))
	}
	return
}

// RandomTextGetByID ...
func RandomTextGetByID(repository *repositories.RandomText, id uint) (dto dtos.RandomText, err error) {
	entity, err := repository.GetByID(id)
	dto = randomTextToDTO(entity)
	return
}
