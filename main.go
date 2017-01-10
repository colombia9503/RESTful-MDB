package main

import (
	"log"
	"net/http"

	"github.com/colombia9503/RESTful-MDB/common"
	"github.com/colombia9503/RESTful-MDB/routers"
	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Escuchando el puerto 8000..")
	http.ListenAndServe(":8000", n)
}
