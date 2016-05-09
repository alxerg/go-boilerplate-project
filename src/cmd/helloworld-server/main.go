package main

import (
	"fmt"
	"helloworld"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/:name", index)

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if name == "" {
		name = "Unknown"
	}
	fmt.Fprint(w, helloworld.Hello(name))
}
