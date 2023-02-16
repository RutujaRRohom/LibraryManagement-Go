package db
import
(
	"context"
	logger "github.com/sirupsen/logrus"
	"library_management/domain"
	"fmt"
	"database/sql"
	//"log"

)

// type Users struct{
// 	 User_id int  `db:"user_id" json:"user_id"`
// 	 Email string  `db:"email" json:"email"`
// 	 Password string `db:"Password" json:"Password"`
// 	 Name string     `db:"Name" json:"Name"`	
// 	 Role string    `db:"role" json:"role"`			
// }
type Storer interface{
	CreateUser(ctx context.Context,users domain.UserResponse) (err error)
	LoginUser(context.Context,string,string) (int, error)
}



func (s *pgStore) CreateUser(ctx context.Context,users domain.UserResponse) (err error){
	sqlQuery:=`INSERT INTO "users"(user_id,email,Password,Name,role) VALUES ($1,$2,$3,$4,$5) returning user_id`
	err = s.db.QueryRow(sqlQuery,&users.User_id,&users.Email,&users.Password,&users.Name,&users.Role).Scan(&users.User_id)
	if err!= nil{
		logger.WithField("err",err.Error()).Error("error registering user")
		return err
	}
	return err
}  



func (s *pgStore) LoginUser(ctx context.Context,Email string,Password string) (u_id int ,err error){
	loginQuery := "SELECT user_id from users where email=$1 and Password=$2"
	err= s.db.QueryRow(loginQuery,Email,Password).Scan(&u_id)	
	// log.Println(s.db.QueryRow("SELECT * from users"))
	// return "hi",nil
	if err != nil {
		if err == sql.ErrNoRows {
			// Email and password combination not found
			return 0, fmt.Errorf("invalid email or password")
		}
		// Other error occurred
		logger.WithField("err", err.Error()).Error("Error during login")
		return 0, err
		// logger.WithField("err", err.Error()).Error("Error incorrect email or password")
		// return
	}
	return u_id,nil

}