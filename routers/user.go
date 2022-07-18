package routers

import (
	"RESTAPI/controllers"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/user", controllers.GetUser)

	fmt.Println("starting web server at http://localhost:8080/")
}
