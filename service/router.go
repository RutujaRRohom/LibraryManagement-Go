package service

import (
	//"fmt"
	"net/http"

	//"library_management/config"

	"github.com/gorilla/mux"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//router.HandleFunc("/users", listUsersHandler(deps)).Methods(http.MethodGet)
	router.HandleFunc("/register", registerUserHandler(deps)).Methods(http.MethodPost)

	router.HandleFunc("/login", loginUserHandler(deps)).Methods(http.MethodPost)

	router.HandleFunc("/addbook", ValidateAdmin(addBooksHandler(deps))).Methods(http.MethodPost)

	router.HandleFunc("/books", getAllBooksHandler(deps)).Methods(http.MethodGet)

	router.HandleFunc("/issue", ValidateUser(issueBookHandler(deps))).Methods(http.MethodPost)

	router.HandleFunc("/updatepassword", ValidateEmail(ResetPasswordHandler(deps))).Methods(http.MethodPost)

	router.HandleFunc("/updatename", ValidateEmail(UpdateNameHandler(deps))).Methods(http.MethodPost)

	router.HandleFunc("/users/email/name", ValidateAdmin(getUserByEmailHandler(deps))).Methods(http.MethodGet)

	router.HandleFunc("/books/activity", ValidateAdmin(getBooksActivityHandler(deps))).Methods(http.MethodGet)

	router.HandleFunc("/users/books/activity", ValidateUser(getBookshandler(deps))).Methods(http.MethodGet)

	router.HandleFunc("/users/return", ValidateUser(ReturnBookHandler(deps))).Methods(http.MethodPost)

	return
}
