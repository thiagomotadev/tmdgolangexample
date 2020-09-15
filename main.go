package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/thiagomotadev/tmdgolangbase"
	"github.com/thiagomotadev/tmdgolangexample/repositories"
	"github.com/thiagomotadev/tmdgolangexample/routes"
)

func main() {
	postgresql1Repository := tmdgolangbase.PostgresBaseRepository{}
	err := postgresql1Repository.Connect("DB1")
	if err != nil {
		log.Fatal(err)
	}

	router := tmdgolangbase.Router{MuxRouter: mux.NewRouter()}

	randomTextRepository := repositories.RandomText{Base: &postgresql1Repository}

	routes.SetRandomTextRoutes(&router, &randomTextRepository)

	router.ListenAndServe()
}
