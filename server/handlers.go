package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nlp_text_classifier/types"

	"github.com/gorilla/mux"
)

func HandlerPlaceholder(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("handler %s\n", request.RequestURI)
}

func GetProfilesHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []types.Person
	var status int
	var err error
	var ctx = request.Context()
	vars := mux.Vars(request)
	db_param, ok := vars["db"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "db parameter not defined" }`))
		return
	}
	fmt.Println(vars)
	people, status, err = p.GetProfiles(ctx, db_param)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func GetProfileHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person types.Person
	var status int
	var err error
	var ctx = request.Context()
	vars := mux.Vars(request)
	db_param, ok := vars["db"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "db parameter not defined" }`))
		return
	}
	id_param, ok := vars["id"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "id parameter not defined" }`))
		return
	}
	fmt.Println(vars)
	person, status, err = p.GetProfile(ctx, db_param, id_param)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetDatabasesHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var databases []string
	var status int
	var err error
	var ctx = request.Context()
	// TODO: add filter params
	// vars := mux.Vars(request)

	databases, status, err = p.GetDatabases(ctx)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(databases)
}
