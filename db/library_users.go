package db
import
(
	"context"
	logger "github.com/sirupsen/logrus"
	"library_management/domain"
	"fmt"
	//"database/sql"
	"log"
	"github.com/google/uuid"

)


type Storer interface{
	CreateUser(ctx context.Context,users domain.UserResponse) (err error)
	LoginUser(context.Context,string,string) (string, error)
	AddingBook(ctx context.Context,add domain.AddBook)(err error)
    GetAllBooksFromDb(ctx context.Context) ([]domain.GetAllBooksResponse , error)
 	getBookById(ctx context.Context,BookId string)(Book domain.GetBookById, err error)
	addUserIssuedBook(ctx context.Context,UserId string,BookId string)(err error)
    updateBookStatus(ctx context.Context,book domain.GetBookById)(err error)
	IssuedBook(ctx context.Context,booking domain.IssueBookRequest)(Book domain.IssuedBookResponse,err error)
	UpdatePassword(ctx context.Context,email string,pass domain.ResetPasswordRequest)(err error)
	Updatename(ctx context.Context,email string,name domain.ResetNameRequest)(err error)

}



func (s *pgStore) CreateUser(ctx context.Context,users domain.UserResponse) (err error){
	sqlQuery:=`INSERT INTO users(user_id,email,Password,Name,role) VALUES ($1,$2,$3,$4,$5) returning user_id`
	err = s.db.QueryRow(sqlQuery,&users.User_id,&users.Email,&users.Password,&users.Name,&users.Role).Scan(&users.User_id)
	if err!= nil{
		logger.WithField("err",err.Error()).Error("error registering user")
		return 
	}
	return err
}  



func (s *pgStore) LoginUser(ctx context.Context,Email string,Password string) (role string ,err error){
	loginQuery := "SELECT role from users where email=$1 and password=$2"
	//loginQuery
	err= s.db.QueryRow(loginQuery,Email,Password).Scan(&role)	
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error incorrect email or password")
		return
	}
	return role,nil

}


func(s *pgStore) AddingBook(ctx context.Context,add domain.AddBook)(err error){
	bookAddQuery:= `INSERT INTO books(book_id,book_name,book_author,publisher,quantity,status) VALUES($1,$2,$3,$4,$5,$6) returning book_id`
	err =s.db.QueryRow(bookAddQuery,&add.BookId,&add.BookName,&add.BookAuthor,&add.Publisher,&add.Quantity,&add.Status).Scan(&add.BookId)
	if err!=nil{
		logger.WithField("err",err.Error()).Error("error in adding book")
		return
	}
	return 

}

func(s *pgStore) GetAllBooksFromDb(ctx context.Context) (books []domain.GetAllBooksResponse, err error){

	getBooksQuery:=`SELECT * from books`
	rows,err :=s.db.Query(getBooksQuery)
	if err!=nil{
		logger.WithField("err",err.Error()).Error("error in getting book")
		return
	}
	for rows.Next(){
		var book domain.GetAllBooksResponse
		err=rows.Scan(&book.BookId,&book.BookName,&book.BookAuthor,&book.Publisher,&book.Quantity,&book.Status)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error scanning books")
			return
		}
		books=append(books,book)
	}
	return 

}

func(s *pgStore)IssuedBook(ctx context.Context,booking domain.IssueBookRequest)(books domain.IssuedBookResponse ,err error){

    
    // getBookByIdQuery:=`SELECT * from books where book_id=$1`
	//var Book domain.GetBookById
	Book,err := s.getBookById(ctx,booking.BookId)
	fmt.Println("rutuja IssuedBook 98")
	if err != nil{
		logger.WithField("error occured",err.Error()).Error("error getting book id")
		return
	}
	if (Book.Status=="notAvailable") || (Book.Quantity<=0){
		logger.WithField("err",err.Error()).Error("book is not available ")
		return
	}
	fmt.Println("rutuja IssuedBook 107", Book)
	//TODO getUser()
	err = s.addUserIssuedBook(ctx,booking.UserId,booking.BookId)
    if err != nil {
		logger.WithField("err",err.Error()).Error("error in adding  book user")
         return
    }

		err=s.updateBookStatus(ctx,Book)
		if err !=nil{
			logger.WithField("err",err.Error()).Error("error in updating book")
			return
        }

		issued:= domain.IssuedBookResponse{
			//issue_id : issueID ,
			UserId  : booking.UserId ,
			BookId  : Book.BookId ,
			BookName :Book.BookName ,
    		BookAuthor :Book.BookAuthor,
   			Publisher :Book.Publisher,
			Quantity :Book.Quantity,
			Status :Book.Status,

		}
		books=issued
		return books ,nil
	}







func (s *pgStore) getBookById(ctx context.Context,BookId string)(Book domain.GetBookById,err error){
	
	// 	stmt, err := s.db.Prepare("SELECT * FROM books  WHERE book_id = ?")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
    //  // var Book domain.GetBookById
	//   err = stmt.QueryRow(BookId).Scan(&Book.BookId, &Book.BookName, &Book.BookAuthor, &Book.Publisher, &Book.Quantity,&Book.Status)
	//   if err != nil {
	// 	 // if err == sql.ErrNoRows {
    //       logger.WithField("err",err.Error()).Error("error occured")
	// 	 //}
	// 	  return Book, err
	//   }
	//   return Book, nil
   err=s.db.QueryRow("SELECT * FROM books  WHERE book_id = $1",BookId).Scan(&Book.BookId,&Book.BookName, &Book.BookAuthor, &Book.Publisher, &Book.Quantity,&Book.Status)
   if err != nil {
	return Book, err
} 
	return Book, nil

}





func (s *pgStore)addUserIssuedBook(ctx context.Context,UserId string,BookId string)(err error){

	issueID:= uuid.New()

    // Insert the book issuing record into the database
    _, err = s.db.Exec("INSERT INTO book_activity ( id,user_id, book_id) VALUES ($1,$2,$3)",issueID,UserId,BookId)
    if err != nil {
		logger.WithField("err",err.Error()).Error("error occured while issuing book")
		return 
    }

    return 
}

func (s *pgStore)updateBookStatus(ctx context.Context, book domain.GetBookById)(err error){
	// Update the book quantity in the database
	quantity := book.Quantity -1
	if quantity <=0{
		_, err = s.db.Exec("UPDATE books SET quantity = $1, status=$2 WHERE book_id = $3", quantity, "notavailable",book.BookId)
		if err!=nil{
			logger.WithField("err",err.Error()).Error("error occured")
				return
		}
		
	}else{
		_, err = s.db.Exec("UPDATE books SET quantity = $1 WHERE book_id = $2",quantity, book.BookId)
		if err != nil {
			logger.WithField("err",err.Error()).Error("error occured")
			return 
		}
	}
   
	
	return
}

func (s *pgStore)UpdatePassword(ctx context.Context,email string,pass domain.ResetPasswordRequest)(err error){
	
	//check if the user exists with this email address in database
	 //resetPasswordQuery:= `select email from users where password=$1`
	 err = s.db.QueryRow("select email from users where password= $1",pass.CurrentPassword ).Scan(&email)
	 fmt.Println(pass.CurrentPassword)
     if err != nil {
	 log.Fatal(err)
	 return
	}
	_,err = s.db.Exec("UPDATE users SET password= $1 WHERE email = $2", pass.NewPassword,email)
		if err!=nil{
			logger.WithField("err",err.Error()).Error("error occured")
				return
		}
	return
     

}

func (s *pgStore)Updatename(ctx context.Context,email string,name domain.ResetNameRequest)(err error){
	
	 err = s.db.QueryRow("select email from users where name=$1",name.CurrentName ).Scan(&email)
	 //fmt.Println(pass.CurrentPassword)
     if err != nil {
		logger.WithField("err",err.Error()).Error("error occured")
		 return
	}
	_,err = s.db.Exec("UPDATE users SET name= $1 WHERE email = $2", name.NewName,email)
		if err!=nil{
			logger.WithField("err",err.Error()).Error("error occured")
				return
		}
	return
     

}




