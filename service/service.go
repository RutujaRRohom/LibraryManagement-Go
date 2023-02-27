package service

import(
   "context"
   "library_management/domain"
   "library_management/db"
    "github.com/sirupsen/logrus"
	//"golang.org/x/crypto/bcrypt"
	"errors"
	"crypto/sha256"
	"encoding/base64"
	"time"
	"github.com/dgrijalva/jwt-go"
	//"fmt"


)

type Services interface{
	RegisterUser(ctx context.Context,users domain.Users) (err error)
    Login(ctx context.Context,userAuth domain.LoginRequest)(token string,err error)
	AddBooks(context.Context,domain.AddBook) (domain.AddBookResponse,error)
	GetBooks(ctx context.Context)([]domain.GetAllBooksResponse,error)
	IssueBook(ctx context.Context,issueReq domain.IssueBookRequest)(booked domain.IssuedBookResponse,err error)
    ResetPassword(ctx context.Context,email string,pass domain.ResetPasswordRequest)(err error)
	UpdateName(ctx context.Context,email string,name domain.ResetNameRequest)(err error)
	getUsersByEmailName(ctx context.Context,emailID string,prefix string)(users []domain.GetUsersResponse,err error)
	GetBooksActivity(ctx context.Context)(book []domain.GetBooksActivityResponse,err error)
	Getbooks(ctx context.Context,email string)(book []domain.GetBooksResponse,err error)
	ReturnBook(ctx context.Context,book domain.ReturnBookRequest)(err error)


}

type bookService struct{
	store db.Storer
}

func NewBookService(s db.Storer) Services{
	return &bookService{
			store:s,
	}

}


var secretKey=[]byte("81mohomrajutr")

func GenerateToken(role string,email string) (token string,err error){
		tokenExpirationTime := time.Now().Add(time.Hour * 24)
		tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Role":    role,
			"email": email,
			"exp":     tokenExpirationTime.Unix(),
		})
		token, err = tokenObject.SignedString(secretKey)
		return
	}


func HashPassword(password string) (string){

	hsha := sha256.New()
	hsha.Write([]byte(password))
	hash := base64.URLEncoding.EncodeToString(hsha.Sum(nil))
	logrus.Info(password, " -> ", hash)
	return hash
}

func (b *bookService)RegisterUser(ctx context.Context,user domain.Users) (err error){

	// registerResponse = domain.UserResponse{
	// 	//User_id :user.User_id,
	// 	Email   :user.Email,
	// 	//Password :user.Password,
	// 	Name :user.Name,
	// 	Role  :user.Role,

	// }
	user.Password=HashPassword(user.Password)
	err = b.store.CreateUser(ctx,user)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("error registering farmer")
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\""{
			err = ErrDuplicateEmail
			return
		}
		return

	}
	return 
}


func(b *bookService) Login(ctx context.Context,userAuth domain.LoginRequest) (token string,err error){

    // var u_id int
	userAuth.Password=HashPassword(userAuth.Password)
	// if err!=nil{
	// 	errors.New("encryption failed")
	// 	return
	// }
	var role string
	role,err = b.store.LoginUser(ctx,userAuth.Email,userAuth.Password)
	if err!=nil{
		errors.New("error")
		return
	}
	
  
	token,err =GenerateToken(role,userAuth.Email)
	if err!=nil{
		logrus.WithField("err", err.Error()).Error("error generating jwt token for user")
		return

	}
	return token,nil
	
}



func(b *bookService) AddBooks(ctx context.Context,add domain.AddBook) (bookAdd domain.AddBookResponse,err error){



	bookAdd = domain.AddBookResponse{
		//BookId:add.BookId,
		BookName:add.BookName,
		BookAuthor:add.BookAuthor,
		Publisher:add.Publisher,
		Quantity:add.Quantity,
		Status:add.Status,
	}
	
	bookAdd.BookId,err=b.store.AddingBook(ctx,bookAdd)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error adding book")
		return
	}

	
	return
}


func (b *bookService) GetBooks(ctx context.Context)(books []domain.GetAllBooksResponse,err error){
	books,err =b.store.GetAllBooksFromDb(ctx)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error in fetching books")
		return
	} 
	return
}


func (b *bookService) IssueBook(ctx context.Context,issueReq domain.IssueBookRequest)(booked domain.IssuedBookResponse,err error){
	booked,err=b.store.IssuedBook(ctx,issueReq)
	if err != nil{
		logrus.WithField("err", err.Error()).Error("error in issuing books")
		return
	}
	return
}


func (b *bookService) ResetPassword(ctx context.Context,email string,pass domain.ResetPasswordRequest)(err error){



	pass.CurrentPassword=HashPassword(pass.CurrentPassword)
	pass.NewPassword=HashPassword(pass.NewPassword)


	err = b.store.UpdatePassword(ctx,email,pass)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error in reseting password")
		return
	}
	return
}


func (b *bookService)UpdateName(ctx context.Context,email string,name domain.ResetNameRequest)(err error){


	err = b.store.Updatename(ctx,email,name)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error in reseting password")
		return
	}
	return
}


func (b *bookService)getUsersByEmailName(ctx context.Context,emailID string,prefix string)(users []domain.GetUsersResponse,err error){
	users,err=b.store.GetUsers(ctx,emailID,prefix)
	if err !=nil {
		logrus.WithField("err", err.Error()).Error("error in getting users")
		return
	}
	return
}

func (b *bookService)GetBooksActivity(ctx context.Context)(book []domain.GetBooksActivityResponse,err error){
	book ,err =b.store.GetBookActivity(ctx)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("error in getting books")
		return
	}
	return
}

func (b *bookService)Getbooks(ctx context.Context,email string)(book []domain.GetBooksResponse,err error){
	book,err = b.store.GetUserBooks(ctx,email)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("error in getting books issued")
		return
	}
	return
}

func (b *bookService)ReturnBook(ctx context.Context,book domain.ReturnBookRequest)(err error){
	err=b.store.ReturnBooks(ctx,book)
	if err != nil{
		logrus.WithField("err", err.Error()).Error("error in getting books issued")
		return
	}
	//fmt.Println(quantity)
	return nil
}


