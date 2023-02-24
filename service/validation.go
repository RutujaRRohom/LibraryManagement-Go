package service
import
(
	"regexp"
	"errors"
	"github.com/dgrijalva/jwt-go"
	//"net/http"
	logger "github.com/sirupsen/logrus"


)


func validateEmail(email string) (err error){
	em:=regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	if !em.MatchString(email){
		err=errors.New("invalid email")
		return
	}
	return 
}



func ValidateJWT(tokenString string) ( err error) {
	tokenObject, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	role := string(claims["Role"].(string))
	if !ok || role != "Admin" || role !="superadmin"{
		logger.WithField("err",err.Error()).Error(" user is not admin or superadmin")
		return
	}
	return
}


func ValidateJWTEmail(tokenString string) ( email string,err error) {
	tokenObject, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	email = string(claims["email"].(string))

	return email,nil
}

func ValidateUserJWT(tokenString string) ( err error) {
	tokenObject, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	role := string(claims["Role"].(string))
	if !ok || role != "user" {
		logger.WithField("err",err.Error()).Error(" user is not end user")
		return
	}
	return
}