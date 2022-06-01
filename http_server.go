package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func router() http.Handler {
	r := mux.NewRouter().StrictSlash(false)

	v1Routes(r)

	return r
}

func v1Routes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.Handle(
		"/healthcheck",
		handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(healthcheckHandler)),
	).Methods("GET")

	v1.Handle(
		"/partners",
		handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(searchPartnerHandler)),
	).Methods("GET")

	v1.Handle(
		"/partners/search",
		handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(searchPartnerHandler)),
	).Methods("POST")

	v1.Handle(
		"/partners",
		handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(createPartnerHandler)),
	).Methods("POST")

	v1.Handle(
		"/partners/{partnerId}",
		handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(showPartnerHandler)),
	).Methods("GET")
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Floor Found!"))
}

func searchPartnerHandler(w http.ResponseWriter, r *http.Request) {

}

func createPartnerHandler(w http.ResponseWriter, r *http.Request) {
	// model := createPartnerModel{}
	// if err := createPartner(&model); err != nil {
	// 	w.Write([]byte(fmt.Sprint("Error_message: %s", err.Error)))
	// }
}

// type createPartnerModel struct {
// }

func showPartnerHandler(w http.ResponseWriter, r *http.Request) {

}
