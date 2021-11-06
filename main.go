package main

import (
	"net/http"

	routes "bartekgo.com/golang_practice/router"
)

func main() {

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)

}
