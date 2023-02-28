

CREATE TABLE users(
   user_id  SERIAL NOT NULL  ,
   email varchar(320) UNIQUE NOT NULL ,
  Password varchar(64) NOT NULL,
  Name varchar(255) NOT NULL ,
  role varchar(128) NOT NULL,
  PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS books (
  book_id  SERIAL NOT NULL  ,    
  book_name varchar(128) NOT NULL,
  book_author varchar(128) NOT NULL,
  publisher varchar(128) NOT NULL ,
  quantity int NOT NULL,
  status varchar(128) NOT NULL ,
  PRIMARY KEY (book_id)
);


CREATE TABLE book_activity(
   activity_id  SERIAL NOT NULL ,
	issue_date timestamp(3) NOT NULL  DEFAULT CURRENT_TIMESTAMP(3),
   
   IsReturned BOOLEAN NOT NULL DEFAULT false,
   user_id INT,
	book_id INT,
   PRIMARY KEY(activity_id),
	CONSTRAINT fk_book_activity1
      FOREIGN KEY(user_id) 
	  REFERENCES users(user_id),
   CONSTRAINT fk_book_activity2
      FOREIGN KEY(book_id) 
	  REFERENCES books(book_id)
);

 ALTER TABLE book_activity ADD return_date TIMESTAMP DEFAULT '0001-01-01 00:00:00';

ALTER TABLE book_activity DROP COLUMN return_date;





