CREATE TABLE IF NOT EXISTS users ( ID bigserial not null, name varchar(255), createdAt timestamp, deletedAt timestamp);

CREATE TABLE IF NOT EXISTS todos ( ID bigserial not null, name varchar(255), content varchar(255), userID bigint, createdAt timestamp, deletedAt timestamp, constraint fk_user_id foreign key (userID) references users(id) on delete cascade);