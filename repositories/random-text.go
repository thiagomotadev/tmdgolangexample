package repositories

import (
	"github.com/thiagomotadev/tmdgolangbase"
	"github.com/thiagomotadev/tmdgolangexample/entities"
)

// RandomText ...
type RandomText struct {
	Base *tmdgolangbase.PostgresBaseRepository
}

// Add ...
func (repository RandomText) Add(entity *entities.RandomText) (err error) {
	err = repository.Base.Add(entity)
	return
}

// Update ...
func (repository RandomText) Update(entity *entities.RandomText) (err error) {
	err = repository.Base.Update(entity)
	return
}

// DeleteByID ...
func (repository RandomText) DeleteByID(id uint) (err error) {
	err = repository.Base.DeleteByID(&entities.RandomText{}, id)
	return
}

// GetByID ...
func (repository RandomText) GetByID(id uint) (entity entities.RandomText, err error) {
	err = repository.Base.GetByID(&entity, id)
	return
}

// GetAll ...
func (repository RandomText) GetAll() (entities []entities.RandomText, err error) {
	err = repository.Base.GetAll(&entities)
	return
}
