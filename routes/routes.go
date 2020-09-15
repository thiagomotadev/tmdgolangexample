package routes

import (
	"net/http"

	"github.com/thiagomotadev/tmdgolangbase"
	"github.com/thiagomotadev/tmdgolangexample/commands"
	"github.com/thiagomotadev/tmdgolangexample/dtos"
	"github.com/thiagomotadev/tmdgolangexample/queries"
	"github.com/thiagomotadev/tmdgolangexample/repositories"
)

// SetRandomTextRoutes ...
func SetRandomTextRoutes(router *tmdgolangbase.Router, repository *repositories.RandomText) {
	subrouter := router.Subrouter("/random-text")

	subrouter.AddRouteOnRoot("POST", func(rar tmdgolangbase.RequestAndResponse) (err error) {
		dto := dtos.RandomText{}
		rar.GetRequestData(&dto)
		err = commands.RandomTextAdd(repository, dto)
		if err != nil {
			return
		}
		rar.SetResponseStatus(http.StatusCreated)
		return
	})

	subrouter.AddRouteOnRoot("PUT", func(rar tmdgolangbase.RequestAndResponse) (err error) {
		dto := dtos.RandomText{}
		rar.GetRequestData(&dto)
		err = commands.RandomTextUpdate(repository, dto)
		if err != nil {
			return
		}
		rar.SetResponseStatus(http.StatusOK)
		return
	})

	subrouter.AddRouteOnRoot("GET", func(rar tmdgolangbase.RequestAndResponse) (err error) {
		dtos, err := queries.RandomTextGetAll(repository)
		if err != nil {
			return
		}
		rar.SetResponseData(http.StatusOK, dtos)
		return
	})

	subrouter.AddRouteOnID("DELETE", func(rar tmdgolangbase.RequestAndResponse, id uint) (err error) {
		err = commands.RandomTextDeleteByID(repository, id)
		if err != nil {
			return
		}
		rar.SetResponseStatus(http.StatusOK)
		return
	})

	subrouter.AddRouteOnID("GET", func(rar tmdgolangbase.RequestAndResponse, id uint) (err error) {
		dto, err := queries.RandomTextGetByID(repository, id)
		if err != nil {
			return
		}
		rar.SetResponseData(http.StatusOK, dto)
		return
	})
}
