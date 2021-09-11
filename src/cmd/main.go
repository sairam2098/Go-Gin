package main

import (
	"gogin/src/handlers"
)

var (
	routeObj  *handlers.Router = &handlers.Router{}
	serverObj *handlers.Server = &handlers.Server{}
)

func main() {
	router := routeObj.GetRouter()
	newRouter := routeObj.NewRouter(router)

	err := serverObj.Start(newRouter, ":7999")

	if err != nil {
		panic("Error starting server")
	}
}
