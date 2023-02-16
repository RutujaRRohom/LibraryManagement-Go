

CREATE TABLE users (
  user_id integer,
  email text,
  Password text,
  Name  text,
  role text

);


CREATE TABLE books(
  book_id integer,
  book_name text,
  book_author text,
  publisher text,
  status text
  );

CREATE TABLE book_activity(
  transaction_id int,
  user_id int,
  book_id int,
  transaction_date DATE
)




