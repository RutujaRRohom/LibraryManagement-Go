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


)

type Services interface{
	RegisterUser(ctx context.Context,users domain.Users) (userAdded domain.UserResponse,err error)
    Login(ctx context.Context,userAuth domain.LoginRequest)(token string,err error)
	AddBooks(context.Context,domain.AddBook) (domain.AddBook,error)
	GetBooks(ctx context.Context)([]domain.GetAllBooksResponse,error)
	IssueBook(ctx context.Context,issueReq domain.IssueBookRequest)(booked domain.IssuedBookResponse,err error)

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

func GenerateToken(role string) (token string,err error){
		tokenExpirationTime := time.Now().Add(time.Hour * 24)
		tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Role":    role,
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

func (b *bookService)RegisterUser(ctx context.Context,user domain.Users) (registerResponse domain.UserResponse,err error){

	registerResponse = domain.UserResponse{
		User_id :user.User_id,
		Email   :user.Email,
		Password :user.Password,
		Name :user.Name,
		Role  :user.Role,

	}
	registerResponse.Password=HashPassword(registerResponse.Password)
	err = b.store.CreateUser(ctx,registerResponse)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error registering user")
		return
	}
	return registerResponse,err
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
	
  
	token,err =GenerateToken(role)
	if err!=nil{
		logrus.WithField("err", err.Error()).Error("error generating jwt token for user")
		return

	}
	return token,nil
	
}



func(b *bookService) AddBooks(ctx context.Context,add domain.AddBook) (added domain.AddBook,err error){

	
	err=b.store.AddingBook(ctx,add)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error adding book")
		return
	}
	added=add
	return added,nil
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


