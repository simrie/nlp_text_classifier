package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nlp_text_classifier/profile"
	"nlp_text_classifier/types"

	"github.com/gorilla/mux"
)

/*
HandlerPlaceholder is a placeholder
*/
func HandlerPlaceholder(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("handler %s\n", request.RequestURI)
}

/*
GetProfilesHandler returns all the profiles in a database
*/
func GetProfilesHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []types.Person
	var status int
	var err error
	var ctx = request.Context()
	vars := mux.Vars(request)
	dbParam, ok := vars["db"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "db parameter not defined" }`))
		return
	}
	fmt.Println(vars)
	people, status, err = p.GetProfiles(ctx, dbParam)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

/*
GetProfileHandler returns a profile given a database and id
*/
func GetProfileHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person types.Person
	var status int
	var err error
	var ctx = request.Context()
	vars := mux.Vars(request)
	dbParam, ok := vars["db"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "db parameter not defined" }`))
		return
	}
	idParam, ok := vars["id"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "id parameter not defined" }`))
		return
	}
	fmt.Println(vars)
	person, status, err = p.GetProfile(ctx, dbParam, idParam)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

/*
GetDatabasesHandler returns a list of databases
*/
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

/*
StoreProfilesHandler stores documents
*/
func StoreProfilesHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")

	var people []types.Person
	var status int
	var stored int
	var err error
	var ctx = request.Context()

	vars := mux.Vars(request)
	dbParam, ok := vars["db"]
	if !ok {
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "db parameter not defined" }`))
		return
	}

	// Pull docs to store from request body. See "an improved handler"
	// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	request.Body = http.MaxBytesReader(response, request.Body, 1048576)

	decoder := json.NewDecoder(request.Body)
	// DisallowUnknownFields errors if user includes unknown fields.
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&people)
	if err != nil {
		status = http.StatusBadRequest
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// convert structured body to interface
	docs := make([]interface{}, len(people))
	for i, v := range people {
		docs[i] = v
	}

	stored, status, err = p.StoreProfiles(ctx, dbParam, docs)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(stored)
}

/*
GetProfileRawDocHandler extracts a profile from a rawDoc
*/
func GetProfileRawDocHandler(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")

	var rawDoc profile.RawDoc
	var status int
	var profile profile.Profile
	var err error
	//var ctx = request.Context()

	// Pull docs to store from request body. See "an improved handler"
	// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

	// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	request.Body = http.MaxBytesReader(response, request.Body, 1048576)
	decoder := json.NewDecoder(request.Body)
	// DisallowUnknownFields errors if user includes unknown fields.
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&rawDoc)
	if err != nil {
		status = http.StatusBadRequest
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	profile, err = rawDoc.TextProfiler()
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(profile)
}
