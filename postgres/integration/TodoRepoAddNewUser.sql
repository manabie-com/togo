INSERT INTO users(
	user_name,
	limited_per_day,
	create_date,
	create_user)
VALUES(
	$1,
	$2,
	NOW(),
	'admin')
RETURNING id