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
BookName string `db:"book_name" json:"book_name"`
BookAuthor string `db:"book_author" json:"bookAuthor"`
Publisher string  `db:"publisher" json:"publisher"`
Quantity int `'db:"quantity" json:"quantity"`
Status string `db:"status" json:"status"`
}

type GetAllBooksResponse struct{
BookId int `json:"bookId"`
BookName string `json:"book_name"`
BookAuthor string `json:"bookAuthor"`
Publisher string  `json:"publisher"`
Quantity int ` json:"quantity"`
Status string `json:"status"`
}

type IssueBookRequest struct{
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}

type IssuedBookResponse struct{
	Transaction_id int `json:"transaction_id"`
	UserId int `json:"user_id`
	BookId int `json:"book_id"`
	BookName string `json:"book_name"`
    BookAuthor string `json:"bookAuthor"`
    Publisher string  `json:"publisher"`
	Quantity int  `json:"quantity"`
	Status string `json:"status"`

}

// {
//     "id": "123456",
//     "title": "The Catcher in the Rye",
//     "author": "J.D. Salinger",
//     "status": "issued",
//     "quantity": 4,
//     "issued_to": "987654",
//     "issued_date": "2023-02-18T10:00:00Z"
// }
