package server

import (
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

func HandlerSelector(p types.DB_Pool, endpoint string) http.HandlerFunc {
	var fn http.HandlerFunc
	switch endpoint {
	case "/databases":
		fn = func(w http.ResponseWriter, r *http.Request) {
			GetDatabasesHandler(p, w, r)
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

func StartRouter(db_pool db_mongo.Pool) {
	router := mux.NewRouter()
	router.HandleFunc("/load/csv/db/{db}/key/{keycol}/text/{textcol}/tag/{tagcol}", HandlerSelector(db_pool, "/load/csv/db")).Methods("POST")
	router.HandleFunc("/load/doc/db/{db}", HandlerSelector(db_pool, "/load/doc/db")).Methods("POST")
	router.HandleFunc("/profile", HandlerSelector(db_pool, "/profile")).Methods("POST")
	router.HandleFunc("/databases", HandlerSelector(db_pool, "/databases")).Methods("GET")
	router.HandleFunc("/profiles/db/{db}", HandlerSelector(db_pool, "/profiles/db/")).Methods("GET")
	router.HandleFunc("/profile/db/{db}/id/{id}", HandlerSelector(db_pool, "/profile/db/id/")).Methods("GET")
	http.ListenAndServe(":12345", router)
}
