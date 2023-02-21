package db
import
(
	"context"
	logger "github.com/sirupsen/logrus"
	"library_management/domain"
	//"fmt"
	"database/sql"
	"log"
	//"github.com/google/uuid"

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
	LoginUser(context.Context,string,string) (string, error)
	AddingBook(ctx context.Context,add domain.AddBook)(err error)
    GetAllBooksFromDb(ctx context.Context) ([]domain.GetAllBooksResponse , error)
 	getBookById(ctx context.Context,BookId int)(Book domain.IssuedBookResponse, err error)
	addUserIssuedBook(ctx context.Context,UserId int,BookId int)(err error)
    updateBookStatus(ctx context.Context,BookId int)(err error)
	IssuedBook(ctx context.Context,booking domain.IssueBookRequest)(Book domain.IssuedBookResponse,err error)



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
	logger.Info(Email,Password)
	// log.Println(s.db.QueryRow("SELECT * from users"))
	// return "hi",nil
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

func(s *pgStore)IssuedBook(ctx context.Context,booking domain.IssueBookRequest)(Book domain.IssuedBookResponse ,err error){

    
    // getBookByIdQuery:=`SELECT * from books where book_id=$1`
	//var Book domain.IssuedBookResponse
	Book,err= s.getBookById(ctx,booking.BookId)
	if err != nil{
		logger.WithField("error occured",err.Error()).Error("error getting book id")
		return
	}
	if (Book.Status=="notAvailable") || (Book.Quantity<=0){
		logger.WithField("err",err.Error()).Error("book is not available ")
		return
	}
	// err=updateBookQuantity(ctx,Book.BookId)
	// if err!=nil{
	// 	logger.WithField("err",err.Error()).Error("error in updating book status")
	// 	return
	// }
	//TODO getUser()
	err = s.addUserIssuedBook(ctx,booking.UserId, Book.BookId)
    if err != nil {
		logger.WithField("err",err.Error()).Error("error in adding  book user")
         return
    }

	if Book.Quantity<=0{
		err=s.updateBookStatus(ctx,Book.BookId)
		if err !=nil{
			logger.WithField("err",err.Error()).Error("error in updating book")
        }
	}
return Book,nil

}




func (s *pgStore) getBookById(ctx context.Context,BookId int)(Book domain.IssuedBookResponse,err error){
	//var Book domain.IssuedBookResponse
	    //  getBookByIdQuery:=`SELECT * from books where book_id=$1`
		//  err:=s.db.QueryRow(getBookByIdQuery,BookId).Scan(&Book)
		//  if err!=nil{
		// 	logger.WithField("err",err.Error()).Error("error in getting book")
		// 	return
		// }
		// book:=rows{}
		// err = rows.Scan(&booked.BookId, &booked.BookName, &booked.BookAuthor, &booked.Publisher, &booked.Quantity)
		// if err != nil {
		// 	if err == sql.ErrNoRows {
		// 		// If no rows were returned, return an error indicating that the book was not found
		// 		return Book{}, fmt.Errorf("book with ID %s not found", id)
		// 	}
		// 	// If there was an error scanning the row, return the error
		// 	return Book{}, err
		// }
		// return Book,nil
		stmt, err := s.db.Prepare("SELECT * FROM books  WHERE book_id = ?")
		if err != nil {
			log.Fatal(err)
		}
      var book domain.IssuedBookResponse
	  err = stmt.QueryRow(BookId).Scan(&book.BookId, &book.BookName, &book.BookAuthor, &book.Publisher, &book.Quantity)
	  if err != nil {
		  if err == sql.ErrNoRows {
			  // Handle the case of no rows returned.
		  }
		  return book, err
	  }
	  return book, nil
}

// func(s *pgStore)updateBookQuantity(ctx context.Context,bookId int)(err error){
// 	// updateBookQuery:=`UPDATE books SET Quantity=Quantity-1 WHERE book_id =$1`
// 	// err=s.db.QueryRow(ctx,updateBookQuery,)
// 	stmt, err := db.Prepare("UPDATE books SET quantity =quantity-1  WHERE id =$1")
//     if err != nil {
//         return err
//     }
//     defer stmt.Close()

//     // Execute the SQL statement with the quantity and book ID parameters
//     _, err = stmt.Exec(bookId)
//     if err != nil {
//         return err
//     }
// 	return nil
// }

func (s *pgStore)addUserIssuedBook(ctx context.Context,UserId int,BookId int)(err error){
	// addUserIssueBookQuery:=`INSERT INTO book_acivity(user_id,book_id) VALUES ($1,$2) returning transaction_id`
	// err = s.db.QueryRow(addUserIssueBookQuery,UserId,BookId).Scan(&UserId)
	// if err!= nil{
	// 	logger.WithField("err",err.Error()).Error("error adding  user")
	// 	return 
	// }
	// return err

	issuingID := uuid.New()

    // Insert the book issuing record into the database
    _, err = s.db.Exec("INSERT INTO book_activity ( issue_id,user_id, book_id,issue_date) VALUES (?,?,?, ?)",issuingID,UserId, BookId,)
    if err != nil {
        return err
    }

    // Update the book quantity in the database
    _, err = s.db.Exec("UPDATE books SET quantity = quantity - 1 WHERE id = ?", BookId)
    if err != nil {
        return err
    }

    return nil
}

func (s *pgStore)updateBookStatus(ctx context.Context,BookId int)(err error){
	_,err = s.db.Exec("UPDATE books SET status = 'notAvailable' WHERE id = ?", BookId)
    if err != nil {
        return err
    }
	return
}