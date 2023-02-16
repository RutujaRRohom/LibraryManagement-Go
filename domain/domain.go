package domain

import
(

)

type Users struct{
	User_id int  `db:"user_id" json:"user_id"`
	Email string  `db:"email" json:"email"`
	Password string `db:"Password" json:"Password"`
	Name string     `db:"Name" json:"Name"`	
	Role string    `db:"role" json:"role"`			
}


type UserResponse struct{
	User_id int  `db:"user_id" json:"user_id"`
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