package api

import (
	"fmt"
	"net/http"

	"github.com/LettsDoSomeCoding/ranko/logging"
	"github.com/gorilla/mux"
)

var homePageGreeting string

// StartServer starts the Ranko API server, which will run indefinitely
func StartServer(greeting string) {
	fmt.Println("Starting server")

	homePageGreeting = greeting
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/player", createPlayerHandler).Methods("POST")
	router.HandleFunc("/players", returnAllPlayers).Methods("GET")

	logging.GetLogger().Fatal(http.ListenAndServe(":8443", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, homePageGreeting)
}

func logHTTPError(w http.ResponseWriter, err error, httpStatus int) {
	logging.GetLogger().Print(err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}
