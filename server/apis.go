package server

import (
	"net/http"

	"nlp_text_classifier/db"

	"github.com/gorilla/mux"
)

/*
HandlerSelector ties the api endspoint to a function in a package
*/
func HandlerSelector(p db.Pool, endpoint string) http.HandlerFunc {
	var fn http.HandlerFunc
	switch endpoint {
	case "/databases":
		fn = func(w http.ResponseWriter, r *http.Request) {
			GetDatabasesHandler(p, w, r)
		}
	case "/profile":
		fn = func(w http.ResponseWriter, r *http.Request) {
			GetProfileRawDocHandler(p, w, r)
		}
	case "/profiles/db/":
		fn = func(w http.ResponseWriter, r *http.Request) {
			GetProfilesHandler(p, w, r)
		}
	case "/profile/db/id/":
		fn = func(w http.ResponseWriter, r *http.Request) {
			GetProfileHandler(p, w, r)
		}
	case "/load/doc/db":
		fn = func(w http.ResponseWriter, r *http.Request) {
			StoreProfilesHandler(p, w, r)
		}
	default:
		fn = func(w http.ResponseWriter, r *http.Request) {
			HandlerPlaceholder(w, r)
		}
	}
	return http.HandlerFunc(fn)
}

/*
StartRouter defines the endpoints that use the db_pool for database connections
*/
func StartRouter(dbPool db.Pool) {
	router := mux.NewRouter()
	router.HandleFunc("/load/csv/db/{db}/key/{keycol}/text/{textcol}/tag/{tagcol}", HandlerSelector(dbPool, "/load/csv/db")).Methods("POST")
	router.HandleFunc("/load/doc/db/{db}", HandlerSelector(dbPool, "/load/doc/db")).Methods("POST")
	router.HandleFunc("/profile", HandlerSelector(dbPool, "/profile")).Methods("POST")
	router.HandleFunc("/databases", HandlerSelector(dbPool, "/databases")).Methods("GET")
	router.HandleFunc("/profiles/db/{db}", HandlerSelector(dbPool, "/profiles/db/")).Methods("GET")
	router.HandleFunc("/profile/db/{db}/id/{id}", HandlerSelector(dbPool, "/profile/db/id/")).Methods("GET")
	http.ListenAndServe(":12345", router)
}
