package models

import "RESTAPI/entities"

var Data = []entities.User{{"E001", "ethan", "ddd", 21}, {"W001", "wick", "sddsds", 22}, {"B001", "bourne", "dsds", 23},
	{"B002", "bond", "dskjd", 23}}

func GetUserById(id string) entities.User {
	var user entities.User
	for _, val := range Data {
		if val.ID == id {
			user = val
			break
		}
	}
	return user
}

// func UpdateUserByid( string) entities.User {
// 	var user entities.User
// 	for _, val := range Data {
// 		if val.ID == id {
// 			user = val
// 			break
// 		}
// 	}
// 	return user
// }
