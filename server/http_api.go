package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nlp_text_classifier/db_mongo"
	"nlp_text_classifier/types"

	"github.com/gorilla/mux"
)

func Test(str string) string {
	res := fmt.Sprintf("%s", str)
	return res
}

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("handler %s\n", request.RequestURI)
}

func PoolHandler(p *db_mongo.Pool) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		GetPeopleEndpoint(p, w, r)
	}
	return http.HandlerFunc(fn)
}

func GetPeopleEndpoint(p *db_mongo.Pool, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []types.Person
	var status int
	var err error
	people, status, err = db_mongo.GetPeople(p)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func StartRouter(db_pool *db_mongo.Pool) {
	router := mux.NewRouter()
	router.HandleFunc("/person", handler).Methods("POST")
	router.HandleFunc("/people", PoolHandler(db_pool)).Methods("GET")
	router.HandleFunc("/person/{id}", handler).Methods("GET")
	http.ListenAndServe(":12345", router)
}
