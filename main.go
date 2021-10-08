package main

import (
	"api/infrastructure/api/handler"
	"api/infrastructure/api/router"
	"api/infrastructure/datastore"
	"api/interface/controllers"
	"api/usecase/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := datastore.NewMySQL()
	r := httprouter.New()
	s := service.NewUserService(db)
	c := controllers.NewUserController(s)
	h := handler.NewUserHandler(c)
	router.NewRouter(r, h)

	defer db.Close()
}