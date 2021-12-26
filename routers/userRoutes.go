package routers

import (
	"encoding/json"
	"fmt"
	"go-library/services"
	"io/ioutil"
	"net/http"
)

var userService services.UserService

func CreateUser(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		var data map[string]string
		json.Unmarshal(reqBody, &data)
		result, err := userService.FindOrCreateUser(
			data["username"],
			data["plainTextPassword"],
			data["fullName"],
			data["email"],
		)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(fmt.Sprintf(`{
				"user": {
					"id": %v,
					"username": "%v",
					"fullName": "%v",
					"email": "%v",
					"createdAt": "%v",
					"updatedAt": "%v"
				}
			}`, result.ID, result.Username, result.FullName, result.Email, result.CreatedAt, result.UpdatedAt)))
		}
	}
}