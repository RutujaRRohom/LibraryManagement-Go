package service

import (
	"library_management/db"
)
// type Dependencies struct {
// 	Store db.Storer
// 	// define other service dependencies
// }


type Dependencies struct{
	bookService Services
}



func InitDependencies()  (deps Dependencies,err error){
	store,err:=db.Init()
	if err !=nil{
		return 
	}

	UserService := NewBookService(store)

	deps=Dependencies{
		bookService:UserService,
	}
	return deps,nil

		
}