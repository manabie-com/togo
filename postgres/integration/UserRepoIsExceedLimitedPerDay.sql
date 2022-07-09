SELECT (u.limited_per_day <= d.count_todos) is_exceed
  FROM users u
  	   INNER JOIN (SELECT d.user_id, COUNT(*) count_todos
					 FROM todos d
				 GROUP BY d.user_id) d
	   ON u.id = d.user_id
 WHERE 1=1
   AND u.id = $1