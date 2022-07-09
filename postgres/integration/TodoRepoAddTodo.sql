INSERT INTO todos(
	todo_id,
	user_id,
	description,
	create_date,
	create_user)
VALUES(
	(SELECT (COUNT(*) + 1) FROM todos WHERE user_id = $1),
	$1,
	$2,
	NOW(),
	'admin')
RETURNING todo_id