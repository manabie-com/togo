
-- Create User, Todos tables --
--> START
DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id serial,
  user_name VARCHAR(25) NOT NULL,
  limited_per_day int NOT NULL,
  create_date timestamp NOT NULL,
  create_user VARCHAR(25) NOT NULL,
  PRIMARY KEY (id)
);

/*
 one to many: User has many Todos
*/

DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
  todo_id int NOT NULL,
  user_id int NOT NULL,
  description VARCHAR(100),
  create_date timestamp NOT NULL,
  create_user VARCHAR(25) NOT NULL,
  PRIMARY KEY (todo_id, user_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
--< END

----------------------------------------------------------

select * from users;
select * from todos;

-- Retrieve inserted todos
SELECT u.id, u.user_name, u.limited_per_day, STRING_AGG (t.description, ', '), t.create_date::date, NOW()::date as today
  FROM users u
  	  ,todos t
 WHERE 1=1
   AND u.id = t.user_id
GROUP BY u.id, u.user_name, u.limited_per_day, t.create_date::date
ORDER BY u.id
;

-- Check IsExceed
SELECT (COUNT(*) >= 7) is_exceed 
  FROM todos 
 WHERE 1=1
   AND user_id = 2 
   AND create_date >= NOW()::date