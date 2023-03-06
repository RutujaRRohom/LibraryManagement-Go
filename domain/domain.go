package domain

type Users struct {
	User_id  int    `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"Password" json:"password"`
	Name     string `db:"Name" json:"name"`
	Role     string `db:"role" json:"role"`
}

type UserResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"msg"`
	Token   string `json:"token"`
}

type AddBook struct {
	// BookID string `db:"book_id" json:"book_id"`
	BookName   string ` json:"book_name"`
	BookAuthor string `json:"book_author"`
	Publisher  string `json:"publisher"`
	Quantity   int    ` json:"quantity"`
	Status     string `json:"status"`
}

type AddBookResponse struct {
	BookID     int    `db:"book_id" json:"book_id"`
	BookName   string `db:"book_name" json:"book_name"`
	BookAuthor string `db:"book_author" json:"book_author"`
	Publisher  string `db:"publisher" json:"publisher"`
	Quantity   int    `'db:"quantity" json:"quantity"`
	Status     string `db:"status" json:"status"`
}

type GetAllBooksResponse struct {
	BookID     string `json:"book_id"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"book_author"`
	Publisher  string `json:"publisher"`
	Quantity   int    ` json:"quantity"`
	Status     string `json:"status"`
}

type GetBookById struct {
	BookID     int    `json:"bookId"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"bookAuthor"`
	Publisher  string `json:"publisher"`
	Quantity   int    ` json:"quantity"`
	Status     string `json:"status"`
}

type IssueBookRequest struct {
	BookID int `json:"book_id"`
}

type IssuedBookResponse struct {
	IssueID    int    `json:"issue_id"`
	UserID     int    `json:"user_id"`
	BookID     int    `json:"bookId"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"bookAuthor"`
	Publisher  string `json:"publisher"`
	IssueDate  string `json:"issue_date"`
}

type GetActivity struct {
	IssueID    int     `json:"issue_id"`
	IssueDate  string  `json:"issue_date"`
	IsReturned bool    `json:"isreturned"`
	UserID     int     `json:"user_id"`
	BookID     int     `json:"book_id"`
	ReturnDate *string `json:"return_date,omitempty"`
}

type ResetPasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type ResetNameRequest struct {
	CurrentName string `json:"current_name"`
	NewName     string `json:"newName"`
}
type ResetNameResponse struct {
	Message string `json:"message"`
}

type GetUsersResponse struct {
	Name       string  `json:"name"`
	BookIssued string  `json:"book_issued"`
	IssueDate  string  `json:"issue_date"`
	ReturnDate *string `json:"return_date,omitempty"`
}
type GetBooksActivityResponse struct {
	BookID     string  `json:"book_id"`
	UserID     string  `json:"user_id"`
	BookName   string  `json:"book_name"`
	UserName   string  `json:"user_name"`
	IssueDate  string  `json:"issue_date"`
	ReturnDate *string `json:"return_date,omitempty"`
}

type GetbooksRequest struct {
	UserID string `json:"user_id"`
}
type GetBooksResponse struct {
	UserName   string  `json:"user_name"`
	BookID     string  `json:"book_id"`
	BookName   string  `json:"book_name"`
	IssueDate  string  `json:"issue_date"`
	ReturnDate *string `json:"return_date,omitempty"`
}

type ReturnBookRequest struct {
	// UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

type ReturnBookResponse struct {
	Message string `json:"message"`
}

type QuantityResponse struct {
	Quantity int `json:"quantity"`
}
