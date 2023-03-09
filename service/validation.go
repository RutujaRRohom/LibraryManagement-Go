package service

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"

	//"net/http"
	logger "github.com/sirupsen/logrus"
)

type key string

const tokenkey key = "token"

func validateEmail(email string) (err error) {
	em := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	if !em.MatchString(email) {
		err = errors.New("invalid email")
		return
	}
	return
}

func ValidateJWT(tokenString string) (role string, err error) {
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
	role = string(claims["Role"].(string))
	if !ok || (role != "Admin" && role != "superadmin") {
		err = errors.New("user is not admin or superadmin")
		logger.WithField("err", err.Error()).Error(" user is not admin or superadmin")
		return
	}
	return role, nil
}

func ValidateJWTEmail(tokenString string) (email string, err error) {
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

	return email, nil
}

func ValidateUserJWT(tokenString string) (err error) {
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
		logger.WithField("err", err.Error()).Error(" user is not end user")
		return
	}
	return
}

func ValidateJWTId(tokenString string) (UserID int, err error) {
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
	UserID = int(claims["user_id"].(float64))

	return UserID, nil
}

func ValidateUser(valid http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		UserID, err := ValidateJWTId(authHeader)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), tokenkey, UserID)
		valid.ServeHTTP(w, req.WithContext(ctx))

	})
}

func ValidateEmail(valid http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		Email, err := ValidateJWTEmail(authHeader)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), tokenkey, Email)
		valid.ServeHTTP(w, req.WithContext(ctx))

	})
}

func ValidateAdmin(valid http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		Role, err := ValidateJWT(authHeader)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), tokenkey, Role)
		valid.ServeHTTP(w, req.WithContext(ctx))

	})
}
