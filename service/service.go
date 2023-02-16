package service

import(
   "context"
   "library_management/domain"
   "library_management/db"
    "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"errors"
	//"crypto/sha256"
	//"encoding/base64"
	"time"
	"github.com/dgrijalva/jwt-go"


)

type Services interface{
	RegisterUser(ctx context.Context,users domain.Users) (userAdded domain.UserResponse,err error)
    Login(ctx context.Context,userAuth domain.LoginRequest)(token string,err error)
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

func GenerateToken(u_id int) (token string,err error){
		tokenExpirationTime := time.Now().Add(time.Hour * 24)
		tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": u_id,
			"exp":     tokenExpirationTime.Unix(),
		})
		token, err = tokenObject.SignedString(secretKey)
		return
	}


func HashPassword(password string) (string,error){
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword),err
}

func (b *bookService)RegisterUser(ctx context.Context,user domain.Users) (registerResponse domain.UserResponse,err error){

	registerResponse = domain.UserResponse{
		User_id :user.User_id,
		Email   :user.Email,
		Password :user.Password,
		Name :user.Name,
		Role  :user.Role,

	}
	registerResponse.Password,err=HashPassword(registerResponse.Password)
	err = b.store.CreateUser(ctx,registerResponse)
	if err !=nil{
		logrus.WithField("err", err.Error()).Error("error registering user")
		return
	}
	return registerResponse,err
}


func(b *bookService) Login(ctx context.Context,userAuth domain.LoginRequest) (token string,err error){

    // var u_id int
	userAuth.Password,err=HashPassword(userAuth.Password)
	if err!=nil{
		errors.New("encryption failed")
		return
	}
	var u_id int
	u_id,err = b.store.LoginUser(ctx,userAuth.Email,userAuth.Password)
	if err!=nil{
		errors.New("error")
		return
	}
	

	token,err =GenerateToken(u_id)
	if err!=nil{
		logrus.WithField("err", err.Error()).Error("error generating fwt token for user")
		return

	}
	return token,nil
	
	


}




