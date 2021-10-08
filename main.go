package main

import (
	"api/config"
	"api/infrastructure/api/router"
	"api/infrastructure/datastore"
	"api/registry"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	config.LoadConfig()

	db := datastore.NewMySQL()
	r := httprouter.New()
	v := validator.New()

	rg :=registry.NewInteractor(db, v)
	h := rg.NewAppHandler()

	router.NewRouter(r, h)

	defer db.Close()

	http.ListenAndServe(":8080", r)
}