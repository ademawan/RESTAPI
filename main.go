package main

import (
	"RESTAPI/routers"
	"net/http"
)

func main() {
	routers.Router()

	http.ListenAndServe(":8080", nil)

}
