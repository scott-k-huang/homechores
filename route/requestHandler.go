package route

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/*
https://tutorialedge.net/golang/creating-restful-api-with-golang/
*/
func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
