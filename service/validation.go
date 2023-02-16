package service
import
(
	"regexp"
	"errors"
)


func validateEmail(email string) (err error){
	em:=regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	if !em.MatchString(email){
		err=errors.New("invalid email")
		return
	}
	return 
}