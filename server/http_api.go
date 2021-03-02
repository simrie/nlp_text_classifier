package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"nlp_text_classifier/db/db_mongo"
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

func PoolHandler(p types.DB_Pool) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		GetPeopleEndpoint(p, w, r)
	}
	return http.HandlerFunc(fn)
}

func GetPeopleEndpoint(p types.DB_Pool, response http.ResponseWriter, request *http.Request) {
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
	people, status, err = p.GetPeople(ctx, db_param)
	if err != nil {
		response.WriteHeader(status)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func StartRouter(db_pool db_mongo.Pool) {
	router := mux.NewRouter()
	router.HandleFunc("/person", handler).Methods("POST")
	router.HandleFunc("/people", PoolHandler(db_pool)).Methods("GET")
	router.HandleFunc("/db/{db}/people", PoolHandler(db_pool)).Methods("GET")
	router.HandleFunc("/person/{id}", handler).Methods("GET")
	http.ListenAndServe(":12345", router)
}
