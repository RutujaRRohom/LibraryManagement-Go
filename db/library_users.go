package db
import
(
	"context"
	logger "github.com/sirupsen/logrus"
	"library_management/domain"
	"fmt"
	"time"
	"log"
	//"github.com/google/uuid"
	"errors"

)


type Storer interface{
	CreateUser(ctx context.Context,users domain.Users) (err error)
	LoginUser(context.Context,string,string) (string, error)
	AddingBook(ctx context.Context,add domain.AddBookResponse)(bookId int,err error)
    GetAllBooksFromDb(ctx context.Context) ([]domain.GetAllBooksResponse , error)
 	getBookById(ctx context.Context,BookID int)(Book domain.GetBookById, err error)
	addUserIssuedBook(ctx context.Context,UserID int,BookID int)(err error)
    updateBookStatus(ctx context.Context,book domain.GetBookById)(err error)
	IssuedBook(ctx context.Context,booking domain.IssueBookRequest)(Book domain.IssuedBookResponse,err error)
	UpdatePassword(ctx context.Context,email string,pass domain.ResetPasswordRequest)(err error)
	Updatename(ctx context.Context,email string,name domain.ResetNameRequest)(err error)
	GetUsers(ctx context.Context,emailID string,prefix string)(users []domain.GetUsersResponse,err error)
	GetBookActivity(ctx context.Context) (book []domain.GetBooksActivityResponse,err error)
	GetUserBooks(ctx context.Context,email string)(book []domain.GetBooksResponse,err error)
	ReturnBooks(ctx context.Context,book domain.ReturnBookRequest)(err error)

}



func (s *pgStore) CreateUser(ctx context.Context,users domain.Users) (err error){
	sqlQuery:=`INSERT INTO users(email,Password,Name,role) VALUES ($1,$2,$3,$4) returning user_id`
	err = s.db.QueryRow(sqlQuery,&users.Email,&users.Password,&users.Name,&users.Role).Scan(&users.User_id)
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


func(s *pgStore) AddingBook(ctx context.Context,bookAdd domain.AddBookResponse)( BookID int,err error){
	fmt.Println("add 62",bookAdd)
	bookAddQuery:= `INSERT INTO books(book_name,book_author,publisher,quantity,status) VALUES($1,$2,$3,$4,$5) returning book_id`
	err =s.db.QueryRow(bookAddQuery,&bookAdd.BookName,&bookAdd.BookAuthor,&bookAdd.Publisher,&bookAdd.Quantity,&bookAdd.Status).Scan(&BookID)
	if err!=nil{
		logger.WithField("err",err.Error()).Error("error in adding book")
		return
	}
	return BookID,nil

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
		err=rows.Scan(&book.BookID,&book.BookName,&book.BookAuthor,&book.Publisher,&book.Quantity,&book.Status)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error scanning books")
			return
		}
		books=append(books,book)
	}
	return 

}

func(s *pgStore)IssuedBook(ctx context.Context,booking domain.IssueBookRequest)(books domain.IssuedBookResponse ,err error){

    var userExists,bookExists bool
    err= s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", booking.UserID).Scan(&userExists)
        if err!=nil {
			logger.WithField("err",err.Error()).Error("user not exist")
			//err=errors.New("user with this id not exist")
			return
        }
		fmt.Println("hiiiiiiii  301",userExists)
        err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE book_id = $1)", booking.BookID).Scan(&bookExists)
        if err != nil {
			logger.WithField("err",err.Error()).Error("book with this ID not exist")
			return
        }
		if !userExists {
			err=errors.New("user not exist with this id")

			logger.WithField("err",err.Error()).Error("user not found")
            return
        }
		fmt.Println("line no 328")
        if !bookExists {
			err=errors.New("book is not issued with this id")

			logger.WithField("err",err.Error()).Error("book not found")
             return
        }
	Book,err := s.getBookById(ctx,booking.BookID)
	//fmt.Println("rutuja IssuedBook 98")
	if err != nil{
		//logger.WithField("error occured",err.Error()).Error("error getting book id")
		//errors.New("invalid book id")
		return
	}
	
	if (Book.Status=="notavailable") && (Book.Quantity<=0){
		err =  errors.New("book not available")
		return
	}
	fmt.Println("error at 107")
	//fmt.Println("rutuja IssuedBook 107", Book)
	//TODO getUser()
	err = s.addUserIssuedBook(ctx,booking.UserID,booking.BookID)
    if err != nil {
		// logger.WithField("err",err.Error()).Error("error in adding  book user")
		err=errors.New("user id or book id does not exist")
         return
    }

		err=s.updateBookStatus(ctx,Book)
		if err !=nil{
			logger.WithField("err",err.Error()).Error("error in updating book")
			return
        }

		issued:= domain.IssuedBookResponse{
			//issue_id : issueID ,
			UserID  : booking.UserID ,
			BookID  : Book.BookID ,
			BookName :Book.BookName ,
    		BookAuthor :Book.BookAuthor,
   			Publisher :Book.Publisher,
			//Quantity :Book.Quantity,
			//Status :Book.Status,

		}
		books=issued
		return 
	}







func (s *pgStore) getBookById(ctx context.Context,BookID int)(Book domain.GetBookById,err error){
	
	
   err=s.db.QueryRow("SELECT * FROM books  WHERE book_id = $1",BookID).Scan(&Book.BookID,&Book.BookName, &Book.BookAuthor, &Book.Publisher, &Book.Quantity,&Book.Status)
   if err != nil {
	err=errors.New("invalid book id ")
    return 
} 
	return 

}

func (s *pgStore)addUserIssuedBook(ctx context.Context,UserID int,BookID int)(err error){

	//issueID:= uuid.New()

    // Insert the book issuing record into the database
    _, err = s.db.Exec("INSERT INTO book_activity ( user_id, book_id) VALUES ($1,$2)",UserID,BookID)
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
		_, err = s.db.Exec("UPDATE books SET quantity = $1, status=$2 WHERE book_id = $3", quantity, "notavailable",book.BookID)
		if err!=nil{
			logger.WithField("err",err.Error()).Error("error occured")
				return
		}
		
	}else{
		_, err = s.db.Exec("UPDATE books SET quantity = $1 WHERE book_id = $2",quantity, book.BookID)
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


func (s *pgStore)GetUsers(ctx context.Context,emailID string,prefix string)(users []domain.GetUsersResponse,err error){
	// getUsersQuery:=`select * from users where email LIKE $1 and name LIKE $2`
	// rows,err:=s.db.Query(getUsersQuery,emailID,prefix)
	rows,err:=s.db.Query("select  users.name,books.book_name , book_activity.issue_date,book_activity.return_date from users INNER JOIN  book_activity on users.user_id = book_activity.user_id INNER JOIN books on books.book_id = book_activity.book_id WHERE email LIKE $1 || '%' OR name LIKE $2 || '%' ", emailID, prefix);

    fmt.Println(err)
	if err!=nil{
		logger.WithField("err",err.Error()).Error("error in getting users")
		return
	}
	for rows.Next(){
		var user domain.GetUsersResponse
		err=rows.Scan(&user.Name,&user.BookIssued,&user.IssueDate,&user.ReturnDate)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error scanning books")
			return
		}
		users=append(users,user)
	}
	return users,nil
}

func (s *pgStore)GetBookActivity(ctx context.Context) (book []domain.GetBooksActivityResponse,err error){
	rows,err:=s.db.Query("select  books.book_id ,users.user_id,books.book_name, users.name , book_activity.issue_date,book_activity.return_date from users INNER JOIN  book_activity on users.user_id = book_activity.user_id INNER JOIN books on books.book_id = book_activity.book_id");
	if err!=nil{
		logger.WithField("err",err.Error()).Error("error in getting users")
		return
	}
	//book_activity.issue_date 
	for rows.Next(){
		var books domain.GetBooksActivityResponse
		err=rows.Scan(&books.BookID,&books.UserID,&books.BookName,&books.UserName,&books.IssueDate,&books.ReturnDate)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error scanning books")
			return
		}
		book=append(book,books)
	}
	return book,nil

}

func (s *pgStore)GetUserBooks(ctx context.Context,email string)(book []domain.GetBooksResponse,err error){
	
	
	
	
	
	rows,err:=s.db.Query("select  users.name,books.book_id ,books.book_name , book_activity.issue_date,book_activity.return_date from users INNER JOIN  book_activity on users.user_id = book_activity.user_id INNER JOIN books on books.book_id = book_activity.book_id WHERE users.email=$1",email);
	if err!=nil{
		logger.WithField("err",err.Error()).Error("error in getting users")
		return
	}
	for rows.Next(){
		var books domain.GetBooksResponse
		err=rows.Scan(&books.UserName,&books.BookID,&books.BookName,&books.IssueDate,&books.ReturnDate)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error scanning books")
			return
		}
		book=append(book,books)
	}
	return 

}


func(s *pgStore)ReturnBooks(ctx context.Context,book domain.ReturnBookRequest)(err error){
	//fmt.Println("bookId",book.BookID)
	//fmt.Println("userId",book.UserID)

	var userExists, bookExists, bookIssued bool
     err= s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", book.UserID).Scan(&userExists)
        if err!=nil {
			//logger.WithField("err",err.Error()).Error("user not exist")
			err=errors.New("user with this id not exist")
			return
        }
		fmt.Println("hiiiiiiii  301",userExists)
        err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE book_id = $1)", book.BookID).Scan(&bookExists)
        if err != nil {
			logger.WithField("err",err.Error()).Error("book with this ID not exist")
			return
        }
		fmt.Println("hiiiiiiii 306",bookExists)
        err = s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM book_activity WHERE user_id = $1 AND book_id = $2)", book.UserID, book.BookID).Scan(&bookIssued)
        if err != nil {
		//err=errors.New("invalid user id or book id ")

			logger.WithField("err",err.Error()).Error("this book activity doesnt exists")
			return
        }
		fmt.Println("hiiiiiiii 313",bookIssued)

		if !userExists {
			err=errors.New("user not exist with this id")

			logger.WithField("err",err.Error()).Error("user not found")
            return
        }
		fmt.Println("line no 328")
        if !bookExists {
			err=errors.New("book is not issued with this id")

			logger.WithField("err",err.Error()).Error("book not found")
             return
        }
        if !bookIssued {
			err=errors.New("book or user not exist for this activity")

			logger.WithField("err",err.Error()).Error("book activity  not found")

            return
        }
		//check if book has already returned
		var isReturned bool 

		err=s.db.QueryRow("select isreturned from book_activity where user_id =$1 and book_id=$2",book.UserID, book.BookID).Scan(&isReturned)
		if err != nil {
			logger.WithField("err",err.Error()).Error("book with this ID not exist")
			return
        }
		if isReturned {
			err=errors.New("book already returned")
			return
		}

		_,err = s.db.Exec("UPDATE book_activity SET return_date=$1 WHERE user_id=$2 and book_id=$3", time.Now(),book.UserID, book.BookID)
		//fmt.Println("hello there")
     	 if err != nil {
			logger.WithField("err",err.Error()).Error("error in updating")
			return
        }
		
		//var isReturned bool 
		isReturned=true
		_, err = s.db.Exec("UPDATE book_activity SET isreturned=$1 WHERE user_id=$2 and book_id=$3", isReturned,book.UserID, book.BookID)
		fmt.Println("hello there")
     	 if err != nil {
			logger.WithField("err",err.Error()).Error("error in updating")
			return
        }
		Book,err := s.getBookById(ctx,book.BookID)
		if err != nil{
			logger.WithField("error occured",err.Error()).Error("error getting books")
			return
		}
	
		quantity:=Book.Quantity+1
		_, err = s.db.Exec("UPDATE books SET quantity=$1 WHERE book_id = $2",quantity, book.BookID)
        if err != nil {
			logger.WithField("err",err.Error()).Error("error in update")
			return
        }
		
		if quantity > 0 {
			_,err = s.db.Exec("UPDATE books SET status=$1 WHERE book_id = $2","Available", book.BookID)
			if err != nil {
				logger.WithField("err",err.Error()).Error("error in update")
				return
			}
		}

       return 
}



//rows, err := s.db.Query("SELECT * FROM users WHERE email LIKE $1 || '%' OR name LIKE $2 || '%' ", emailID, prefix)
