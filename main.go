package main

import (
	"RESTAPI/configs"
	"RESTAPI/routers"
	"RESTAPI/utils"
	"fmt"
	"net/http"
)

func main() {
	routers.Router()
	configs := configs.GetConfig()
	db, err := utils.InitDB(configs)
	if err != nil {
		panic(err.Error())
	}
	if db != nil {
		fmt.Println("connected")
	}
	http.ListenAndServe(":8080", nil)

}
