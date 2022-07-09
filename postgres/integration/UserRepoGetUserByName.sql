SELECT id, user_name, limited_per_day
  FROM users
 WHERE 1=1
   AND user_name = $1