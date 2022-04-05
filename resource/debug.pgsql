SELECT
	*
FROM
	users
WHERE
	username = 'NotFound';

SELECT
	*
FROM
	users
WHERE
	username = 'ixcarj';

SELECT
	*
FROM
	tasks
WHERE
	OWNER = 'ixcarj' AND
	deleted = FALSE;

SELECT
	COUNT(created_at)
FROM
	tasks
WHERE
	OWNER = 'tienla' AND
	created_at :: DATE = NOW() :: DATE;

UPDATE
	users
SET
	daily_quantity = 10
WHERE
	username = 'ixcarj' RETURNING username,
	hashed_password,
	full_name,
	email,
	daily_cap,
	daily_quantity,
	password_change_at,
	created_at;
