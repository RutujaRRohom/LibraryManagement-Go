package domain

import
(
// "time"
)

type Users struct{
	User_id string  `db:"user_id" json:"user_id"`
	Email string  `db:"email" json:"email"`
	Password string `db:"Password" json:"Password"`
	Name string     `db:"Name" json:"Name"`	
	Role string    `db:"role" json:"role"`			
}


type UserResponse struct{
	User_id string  `db:"user_id" json:"user_id"`
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
BookId string `db":book_id" json:"bookId"`
BookName string `db:"book_name" json:"book_name"`
BookAuthor string `db:"book_author" json:"bookAuthor"`
Publisher string  `db:"publisher" json:"publisher"`
Quantity int `'db:"quantity" json:"quantity"`
Status string `db:"status" json:"status"`
}

type GetAllBooksResponse struct{
BookId string `json:"bookId"`
BookName string `json:"book_name"`
BookAuthor string `json:"bookAuthor"`
Publisher string  `json:"publisher"`
Quantity int ` json:"quantity"`
Status string `json:"status"`
}

type GetBookById struct{
BookId string `json:"bookId"`
BookName string `json:"book_name"`
BookAuthor string `json:"bookAuthor"`
Publisher string  `json:"publisher"`
Quantity int ` json:"quantity"`
Status string `json:"status"`
}

type IssueBookRequest struct{
	UserId string `json:"user_id"`
	BookId string `json:"bookId"`
}

type IssuedBookResponse struct{
	//issue_id int `json:"issue_id"`
	UserId string `json:"user_id`
	BookId string `json:"bookId"`
	BookName string `json:"book_name"`
    BookAuthor string `json:"bookAuthor"`
    Publisher string  `json:"publisher"`
	Quantity int  `json:"quantity"`
	Status string `json:"status"`
	//Book_Issued_at time.Time ` json:"issue_date"`

}
type issuedBookJson struct{
	message string `json:"message"`
	issued IssuedBookResponse `json:"issued"`
}

type ResetPasswordRequest struct{
	CurrentPassword string `json:"currentPassword"`
	NewPassword string `json:"newPassword"`
}
type ResetPasswordResponse struct{
	Message string `json:"message"`
	}


type ResetNameRequest struct{
	CurrentName string `json:"current_name"`
	NewName string `json:"newName"`
	}
type ResetNameResponse struct{
	Message string `json:"message"`
	}

type GetUsersResponse struct{
	UserID string  ` json:"user_id"`
	Email string  `json:"email"`
	Password string ` json:"Password"`
	Name string     `json:"Name"`	
	Role string    ` json:"role"`	
} 
type GetBooksActivityResponse struct{
	BookID string `json:"book_id"`
	UserID string `json:"user_id"`
	BookName string `json:"book_name"`
	UserName string `json:"user_name"`
	IssueDate string `json:"issue_date"`
}

type GetbooksRequest struct{
   UserID string `json:"user_id"`
}
type GetBooksResponse struct{
	UserName string `json:"user_name"`
	BookID string `json:"book_id"`
	BookName string `json:"book_name"`
	IssueDate string `json:"issue_date"`

}

type ReturnBookRequest struct{
	UserID string `json:"user_id"`
	BookID string `json:"bookId"`
}

type ReturnBookResponse struct{
	Message string `json:"message"`
	}

	