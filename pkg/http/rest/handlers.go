package rest

import (
	"fmt"
	"github.com/Abacode7/delivery-service/pkg/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

func Handlers(service user.Service)  *mux.Router{
	router := mux.NewRouter()

	router.Handle("/api/create/user", createUser(service)).Methods("POST")
	router.Handle("/api/get/users", getAllUser(service)).Methods("GET")
	router.Handle("/api/get/users/{id}", getUser(service)).Methods("GET")

	return router
}

func createUser(service user.Service) http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//Your code here
	})
}

func getAllUser(service user.Service) http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//Your code here
	})
}

func getUser(service user.Service) http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		pathVariables := mux.Vars(request)
		fmt.Println(pathVariables["id"])
		//other code here
	})
}