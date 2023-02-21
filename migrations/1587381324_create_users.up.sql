

-- CREATE TABLE users (
--   user_id integer PRIMARY KEY,
--   email text,
--   Password text,
--   Name  text,
--   role text

-- );


CREATE TABLE books(
  book_id integer PRIMARY KEY,
  book_name text,
  book_author text,
  publisher text,
  quantity integer 
  status text
  );

CREATE TABLE book_activity(
  transaction_id int PRIMARY KEY,
  user_id int FOREIGN KEY references users,
  book_id int FOREIGN KEY references books,
  transaction_date DATE
)




