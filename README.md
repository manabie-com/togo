# Solutions
## Database:
### Prerequisite: postgres 13+
### Config
    user = tasker
    pass = 1qaz2wsx3edc
    host = 127.0.0.1
    db  = tasks
    port  = 5432
### Database & Users: 
    CREATE DATABASE tasks ENCODING UTF8;
    CREATE USER tasker WITH PASSWORD '1qaz2wsx3edc';
    GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO tasker;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO tasker;
    
#### Tables
    -- public.daily_limit definition
    -- Drop table
    -- DROP TABLE public.daily_limit;
    CREATE TABLE public.daily_limit (
        id bigserial NOT NULL,
        user_id bigint NOT NULL,
        task_limit int4 NOT NULL,
        task_date date NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_by text NULL,
        CONSTRAINT daily_limit_pk PRIMARY KEY (id)
    );
    CREATE INDEX daily_limit_user_id_task_date_idx ON public.daily_limit (user_id,task_date);

    -- public.tasks definition
    -- Drop table
    -- DROP TABLE public.tasks;
    CREATE TABLE public.tasks (
    id bigserial NOT NULL,
    summary text NOT NULL,
    description text NULL,
    assignee bigint NOT NULL,
    task_date date NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT tasks_pk PRIMARY KEY (id)
    );
    CREATE INDEX tasks_assignee_task_date_idx ON public.tasks USING btree (assignee, task_date);
#### Function
    CREATE OR REPLACE FUNCTION submit_task(s TEXT, d TEXT, a BIGINT, td DATE) 
    RETURNS BIGINT AS
    $$
    DECLARE
        total_daily_task BIGINT := 0;
        limit_daily_task INT4 := 0;
        id_task BIGINT := 0;
    BEGIN
        PERFORM pg_advisory_lock(a);
        SELECT INTO limit_daily_task SUM(task_limit) FROM daily_limit dl WHERE user_id = a AND task_date = td;
        --RAISE NOTICE 'Limit tasks: %', limit_daily_task;
        SELECT INTO total_daily_task COUNT(*) FROM tasks t WHERE assignee = a AND task_date = td;
        --RAISE NOTICE 'Total tasks: %', total_daily_task;
        IF limit_daily_task > total_daily_task
        THEN
            INSERT INTO tasks (summary, description, assignee, task_date)
            VALUES (s, d, a, td) RETURNING id INTO id_task;
        END IF;
        PERFORM pg_advisory_unlock(a);  
        RETURN id_task;
    END
    $$
        LANGUAGE 'plpgsql'
        RETURNS NULL ON NULL INPUT;
#### Init data
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(1, 5, '2022-05-27', 1);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(1, 6, '2022-05-28', 1);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(1, 7, '2022-05-29', 1);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(1, 8, '2022-05-30', 1);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(1, 9, '2022-05-31', 1);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(2, 15, '2022-05-27', 2);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(2, 16, '2022-05-28', 2);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(2, 17, '2022-05-29', 2);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(2, 18, '2022-05-30', 2);
    INSERT INTO public.daily_limit
    (user_id, task_limit, task_date, updated_by)
    VALUES(2, 19, '2022-05-31', 2);
## Run & Test
  ### Prerequisite: 
    go 1.16
    github.com/beego/beego/v2 v2.0.2
    github.com/smartystreets/goconvey v1.6.4
	gorm.io/driver/postgres v1.3.6
	gorm.io/gorm v1.23.5
  ### Config:
    Copy source code to your GOPATH
  ### Module
    cd $GOPATH/togo
    go mod tidy
  ### Run
    bee run
  ### CURL
    Healthcheck: curl -X GET http://localhost:8080
    Submit task: 
      curl -H "Content-Type: application/json" -X POST -d '{"summary":"Todo task 2022-05-27","description":"do something","assignee":1,"taskDate":"2022-05-27"}' http://localhost:8080/v1/tasks
  ### Testing
  ### Note:
    Before testing, make sure added daily_limit and tasks are empty
  #### Unit test
    cd $GOPATH/togo
    go test unit_test.go -v
  #### Integration test
    cd $GOPATH/togo
    go test integration_test.go -v
      
