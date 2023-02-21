package service 

import
(
"library_management/domain"
"net/http"
"encoding/json"
//"strings"

)

func registerUserHandler(dep Dependencies) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,req *http.Request){
      
		 var user domain.Users
		 err:= json.NewDecoder(req.Body).Decode(&user)
		 if err!=nil{
			http.Error(w,"invalid request",http.StatusBadRequest)
			return
		 }

		 if err=validateEmail(user.Email); err!=nil{
			http.Error(w,"invalid mail",http.StatusBadRequest)
			return
		 }

		 userAdded,err1:= dep.bookService.RegisterUser(req.Context(),user)
		 if err1 != nil{
			http.Error(w,"fail to register user",http.StatusInternalServerError)
			return
		 }
		 json_response,err:=json.Marshal(userAdded)
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
 //TODO validate token
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