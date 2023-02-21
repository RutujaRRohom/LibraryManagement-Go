

CREATE TABLE users(
   user_id varchar(64) NOT NULL,
   email varchar(320) NOT NULL ,
  Password varchar(64) NOT NULL,
  Name varchar(255) NOT NULL ,
  role varchar(128) NOT NULL,
  PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS books (
  book_id varchar(36) NOT NULL ,
  book_name varchar(128) NOT NULL,
  book_author varchar(128) NOT NULL,
  publisher varchar(128) NOT NULL ,
  quantity int NOT NULL,
  status varchar(128) NOT NULL ,
  PRIMARY KEY (book_id)
);


CREATE TABLE book_activity(
   issue_id varchar(32) NOT NULL,
	issue_date timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
   user_id varchar(64),
	book_id varchar(36),
   PRIMARY KEY(issue_id),
	CONSTRAINT fk_book_activity1
      FOREIGN KEY(user_id) 
	  REFERENCES users(user_id),
   CONSTRAINT fk_book_activity2
      FOREIGN KEY(book_id) 
	  REFERENCES books(book_id)
);






