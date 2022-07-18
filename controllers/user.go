package controllers

import (
	"RESTAPI/models"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := models.Data

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		res := models.GetUserById(id)
		result, err = json.Marshal(res)

		// for _, each := range users {
		// 	if each.ID == id {
		// 		result, err = json.Marshal(each)

		// 		if err != nil {
		// 			http.Error(w, err.Error(), http.StatusInternalServerError)
		// 			return
		// 		}

		// 		w.Write(result)
		// 		return
		// 	}
		// }
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
