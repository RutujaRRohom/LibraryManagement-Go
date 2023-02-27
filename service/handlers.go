package service 

import
(
"library_management/domain"
"net/http"
"encoding/json"
"library_management/api"
//"strings"
//"fmt"

)

func registerUserHandler(dep Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
      
		 var user domain.Users
		 err:= json.NewDecoder(req.Body).Decode(&user)
		 if err!=nil{
			http.Error(w,"invalid request",http.StatusBadRequest)
			return
		 }
		 if  user.Email == "" ||  user.Password == "" ||user.Name=="" || user.Role=="" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}


		 if err=validateEmail(user.Email); err!=nil{
			http.Error(w,"invalid mail",http.StatusBadRequest)
			return
		 }

		 err= dep.bookService.RegisterUser(req.Context(),user)
		 if err != nil {
			api.Response(w, http.StatusBadRequest, api.Message{Msg: err.Error()})
			return
		}
		 registerRes:=domain.UserResponse{
			Message:"user register successfully",
			
		 }
		 json_response,err:=json.Marshal(registerRes)
		 if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}
		 w.Header().Add("Content-Type", "application/json")
		 w.Write(json_response)



	})
}



func loginUserHandler(dep Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
     
		 var userAuth domain.LoginRequest
		err := json.NewDecoder(req.Body).Decode(&userAuth)
		if err !=nil{
			http.Error(w,"invalid request",http.StatusBadRequest)
			return
		}

		if  userAuth.Email == "" ||  userAuth.Password == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err=validateEmail(userAuth.Email); err!=nil{
			http.Error(w,"invalid mail",http.StatusBadRequest)
			return
		 }

		 token,err:=dep.bookService.Login(req.Context(),userAuth)
		 if err !=nil{
			http.Error(w,"invalid email or password",http.StatusBadRequest)
			return
		 }

		 loginRes:=domain.LoginResponse{
			Message:"login successful",
			Token  :token,
		 }
		json_response,err:=json.Marshal(loginRes)
		 w.Write(json_response)




	})
}



func addBooksHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
          
		var add domain.AddBook
		err:=json.NewDecoder(req.Body).Decode(&add)
		if err!=nil{
			http.Error(w,"invalid request body",http.StatusBadRequest)
		}
		if add.BookName=="" || add.BookAuthor=="" || add.Publisher=="" || add.Quantity== 0|| add.Status==""{
			http.Error(w,"invalid requets body",http.StatusBadRequest)
		}
 		authHeader := req.Header.Get("Authorization")
		 if authHeader == "" {
		 http.Error(w, "Authorization header is required", http.StatusUnauthorized)
	 		return
 		}
		
		    // Extract the token from the Authorization header
		//tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		err = ValidateJWT(authHeader)
		if err != nil {
            http.Error(w,"invalid token",http.StatusUnauthorized)	
			return
		}
		
		
		//adding books
		addbook,err:=deps.bookService.AddBooks(req.Context(),add)
		if err != nil{
			http.Error(w,"error in adding",http.StatusBadRequest)
		}
		json_response,err:=json.Marshal(addbook)
		if err != nil{
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}
		w.Write(json_response)



	})


}

func getAllBooksHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
		
		books,err:=deps.bookService.GetBooks(req.Context())
		if err!= nil{
			http.Error(w,"error returning books",http.StatusBadRequest)
			return
		}
		json_response,err:=json.Marshal(books)
		if err!=nil{
			http.Error(w,"error in marshelling",http.StatusBadRequest)
			return

		}
		w.Write(json_response)
	})
}


func issueBookHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
        var issueReq domain.IssueBookRequest
		err:=json.NewDecoder(req.Body).Decode(&issueReq)
		if err!=nil{
			http.Error(w,"invalid request body",http.StatusBadRequest)
		}
		if  issueReq.UserId == 0 ||  issueReq.UserId == 0{
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		booked,err:=deps.bookService.IssueBook(req.Context(),issueReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
         	return
		}
		// var response domain.issuedBookJson
		// json_response:= response{
		// 	message:"book issued successfully",
		// 	issued: booked,

		// }
		//message:="book issued successfully"
		bookJSON, err := json.Marshal(booked)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bookJSON)
	})
}



func ResetPasswordHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request ){
		var pass domain.ResetPasswordRequest
		err:=json.NewDecoder(req.Body).Decode(&pass)
		if err !=nil{
			http.Error(w,"invalid request body",http.StatusBadRequest)
		}
	
		
		authHeader := req.Header.Get("Authorization")
		 if authHeader == "" {
		 http.Error(w, "Authorization header is required", http.StatusUnauthorized)
	 		return
 		}
		email,err:=ValidateJWTEmail(authHeader)
		if err != nil{
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		if err=validateEmail(email); err!=nil{
			http.Error(w,"invalid mail",http.StatusBadRequest)
			return
		 }

		err=deps.bookService.ResetPassword(req.Context(),email,pass)
		if err != nil{
			http.Error(w, "invalid current password", http.StatusInternalServerError)
			return
		}
		msg:=domain.ResetPasswordResponse{
			Message:"password reset successfully",
		}
		json_response,err:=json.Marshal(msg)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json_response)
	})
}


func UpdateNameHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request ){
		var name domain.ResetNameRequest
		err:=json.NewDecoder(req.Body).Decode(&name)
		if err !=nil{
			http.Error(w,"invalid request body",http.StatusBadRequest)
		}
	
		
		authHeader := req.Header.Get("Authorization")
		 if authHeader == "" {
		 http.Error(w, "Authorization header is required", http.StatusUnauthorized)
	 		return
 		}
		email,err:=ValidateJWTEmail(authHeader)
		if err != nil{
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		if err=validateEmail(email); err!=nil{
			http.Error(w,"invalid mail",http.StatusBadRequest)
			return
		 }

		err=deps.bookService.UpdateName(req.Context(),email,name)
		if err != nil{
			http.Error(w, "invalid current name", http.StatusInternalServerError)
			return
		}
		msg:=domain.ResetPasswordResponse{
			Message:"name reset successfully",
		}
		json_response,err:=json.Marshal(msg)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json_response)
	})
}


func getUserByEmailNameHandler(deps Dependencies)(http.HandlerFunc){
	return http.HandlerFunc(func (w http.ResponseWriter,req *http.Request){

       // vars:=mux.Vars(req)
	    emailID := req.URL.Query().Get("email_pre")
		namePrefix := req.URL.Query().Get("prefix")
		if emailID == "" || namePrefix=="" {
		 http.Error(w, "query paramter  is required", http.StatusUnauthorized)
	 		return
 		}
    	
		authHeader := req.Header.Get("Authorization")
		 if authHeader == "" {
		 http.Error(w, "Authorization header is required", http.StatusUnauthorized)
	 		return
 		}
		err := ValidateJWT(authHeader)
		if err != nil {
            http.Error(w,"invalid token",http.StatusUnauthorized)	
			return
		}
		
		users,err := deps.bookService.getUsersByEmailName(req.Context(),emailID,namePrefix)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
         	return
		}
		userJSON,err := json.Marshal(users)
		if err!=nil{
			http.Error(w,"error in marshelling",http.StatusBadRequest)
			return

		}
		//w.Header().Set("Content-Type", "application/json")

		w.Write(userJSON)


	})
}


func getBooksActivityHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){

		authHeader := req.Header.Get("Authorization")
		 if authHeader == "" {
		 http.Error(w, "Authorization header is required", http.StatusUnauthorized)
	 		return
 		}
		 err := ValidateJWT(authHeader)
		 if err != nil {
			 http.Error(w,"invalid token",http.StatusUnauthorized)	
			 return
		 }
		 
		book,err:=deps.bookService.GetBooksActivity(req.Context())
		if err!= nil{
			http.Error(w,"error returning books",http.StatusBadRequest)
			return
		}
		json_response,err:=json.Marshal(book)
		if err!=nil{
			http.Error(w,"error in marshelling",http.StatusBadRequest)
			return

		}
		w.Write(json_response)

	})
}


func getBookshandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
	
		authHeader := req.Header.Get("Authorization")
		 if authHeader == "" {
		 http.Error(w, "Authorization header is required", http.StatusUnauthorized)
	 		return
 		}
		 email,err:=ValidateJWTEmail(authHeader)
		 if err != nil{
			 http.Error(w, "invalid token", http.StatusUnauthorized)
			 return
		 }
		 if err=validateEmail(email); err!=nil{
			 http.Error(w,"invalid mail",http.StatusBadRequest)
			 return
		  }
		 book,err:=deps.bookService.Getbooks(req.Context(),email)
		 if err!= nil{
			http.Error(w,"error returning books",http.StatusBadRequest)
			return
		}
		json_response,err:=json.Marshal(book)
		if err!=nil{
			http.Error(w,"error in marshelling",http.StatusBadRequest)
			return

		}
		w.Write(json_response)

	})
}

func ReturnBookHandler(deps Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
		var book domain.ReturnBookRequest
		err:= json.NewDecoder(req.Body).Decode(&book)
		if err!=nil{
			http.Error(w,"invalid request body",http.StatusBadRequest)
			return
		}

		err=deps.bookService.ReturnBook(req.Context(),book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
         	return
		}
		//fmt.Println(status)

		returned:=domain.ReturnBookResponse{
			Message:"book returned successfully",
		}
		bookJSON, err := json.Marshal(returned)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bookJSON)

	})
}
