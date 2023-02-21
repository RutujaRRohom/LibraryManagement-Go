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


// func AdminMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Get the Authorization header from the request
//         authHeader := r.Header.Get("Authorization")
//         if authHeader == "" {
//             http.Error(w, "Authorization header is required", http.StatusUnauthorized)
//             return
//         }

//         // Extract the token from the Authorization header
//         tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

//         // Parse the token and verify its signature
//         token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//             // Replace this with your own secret key or public key depending on how you generate and sign the token
//             return []byte("secret-key"), nil
//         })
//         if err != nil || !token.Valid {
//             http.Error(w, "Invalid token", http.StatusUnauthorized)
//             return
//         }

//         // Extract the user's role from the token's claims
//         claims, ok := token.Claims.(jwt.MapClaims)
//         if !ok {
//             http.Error(w, "Invalid token claims", http.StatusUnauthorized)
//             return
//         }
//         role, ok := claims["role"].(string)
//         if !ok || role != "admin" {
//             http.Error(w, "User is not authorized", http.StatusForbidden)
//             return
//         }

//         // Call the next handler if the user is authorized
//         next.ServeHTTP(w, r)
//     })
// }


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
	if !ok || role != "Admin" {
		logger.WithField("err",err.Error()).Error("error registering user")
		return
	}
	return
}