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
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//router.HandleFunc("/users", listUsersHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("register", registerUserHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("login", loginUserHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/addbook", addBooksHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/books", getAllBooksHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/issue", issueBookHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/updatepassword", ResetPasswordHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/updatename", UpdateNameHandler(deps)).Methods(http.MethodPost)
	router.HandleFunc("/users/email/name", getUserByEmailNameHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/books/activity", getBooksActivityHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/users/books/activity", getBookshandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/users/return", ReturnBookHandler(deps)).Methods(http.MethodPost)
	return
}
