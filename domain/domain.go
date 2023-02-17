package domain

import
(

)

type Users struct{
	User_id int  `db:"user_id" json:"user_id"`
	Email string  `db:"email" json:"email"`
	Password string `db:"Password" json:"Password"`
	Name string     `db:"Name" json:"Name"`	
	Role string    `db:"role" json:"role"`			
}


type UserResponse struct{
	User_id int  `db:"user_id" json:"user_id"`
	Email string  `db:"email" json:"email"`
	Password string `db:"Password" json:"Password"`
	Name string     `db:"Name" json:"Name"`	
	Role string    `db:"role" json:"role"`	

}


type LoginRequest struct{
	Email string `json:"email"`
	Password string `json:"Password"`
}

type LoginResponse struct{
	Message string `json:"msg"`
	Token string `json:"token"`
}


type AddBook struct{
BookId int `db":book_id" json:"bookId"`
BookName string `db:"book_name" json:"book_name`
BookAuthor string `db:"book_author" json:"bookAuthor"`
Publisher string  `db:"publisher" json:"publisher"`
Quantity int `'db:"quantity" json:"quantity`
}

