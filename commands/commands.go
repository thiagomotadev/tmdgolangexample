package commands

import (
	"github.com/thiagomotadev/tmdgolangbase"
	"github.com/thiagomotadev/tmdgolangexample/dtos"
	"github.com/thiagomotadev/tmdgolangexample/entities"
	"github.com/thiagomotadev/tmdgolangexample/repositories"

)

func randomTextToEntity(dto dtos.RandomText) (entity entities.RandomText) {
	entity = entities.RandomText{
		ID:      dto.ID,
		Text:    dto.Text,
		TextSum: tmdgolangbase.SumSHA256(dto.Text),
	}
	return
}

// RandomTextAdd ...
func RandomTextAdd(repository *repositories.RandomText, dto dtos.RandomText) (err error) {
	entity := randomTextToEntity(dto)
	err = repository.Add(&entity)
	if err != nil {
		return
	}

	return
}

// RandomTextUpdate ...
func RandomTextUpdate(repository *repositories.RandomText, dto dtos.RandomText) (err error) {
	entity := randomTextToEntity(dto)
	err = repository.Update(&entity)
	return
}

// RandomTextDeleteByID ...
func RandomTextDeleteByID(repository *repositories.RandomText, id uint) (err error) {
	err = repository.DeleteByID(id)
	if err != nil {
		return
	}

	return
}
