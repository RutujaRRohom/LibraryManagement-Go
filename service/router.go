package service

import (
	//"fmt"
	"net/http"

	//"library_management/config"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	// No version requirement for /ping
	//router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//router.HandleFunc("/users", listUsersHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/registerUser",registerUserHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/loginuser",loginUserHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/addBooks",addBooksHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/getAllBooks",getAllBooksHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/issueBook",issueBookHandler(deps)).Methods(http.MethodPost)
	return
}
