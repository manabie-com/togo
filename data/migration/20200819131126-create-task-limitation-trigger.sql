
-- +migrate Up
-- +migrate StatementBegin
CREATE FUNCTION check_task_limitation() RETURNS trigger AS 
$$
BEGIN
		IF (SELECT count(*) FROM tasks WHERE user_id = NEW.user_id AND created_date = NEW.created_date) >= (select max_todo from users where id=NEW.user_id) THEN
			RAISE EXCEPTION 'reach create_task_limit';
			RETURN NEW;
		ELSE
			RETURN NEW;
		END IF;
END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER create_task_limitation
	BEFORE INSERT ON tasks
	FOR EACH ROW EXECUTE PROCEDURE check_task_limitation();

-- +migrate Down
DROP TRIGGER IF EXISTS create_task_limitation ON "tasks" CASCADE;
DROP FUNCTION IF EXISTS check_task_limitation;