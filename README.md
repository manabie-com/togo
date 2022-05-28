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
        uid text NOT NULL,
        task_limit int4 NOT NULL,
        task_date date NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_by text NULL,
        CONSTRAINT daily_limit_pk PRIMARY KEY (id)
    );
    CREATE INDEX daily_limit_uid_task_date_idx ON public.daily_limit (uid,task_date);

    -- public.tasks definition
    -- Drop table
    -- DROP TABLE public.tasks;
    CREATE TABLE public.tasks (
    id bigserial NOT NULL,
    summary text NOT NULL,
    description text NULL,
    assignee text NOT NULL,
    task_date date NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT tasks_pk PRIMARY KEY (id)
    );
    CREATE INDEX tasks_assignee_task_date_idx ON public.tasks USING btree (assignee, task_date);
#### Init data
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('NYUh5d02ZRYLtxI4QriCw1cz2ux1', 5, '2022-05-27', 'NYUh5d02ZRYLtxI4QriCw1cz2ux1');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('NYUh5d02ZRYLtxI4QriCw1cz2ux1', 6, '2022-05-28', 'NYUh5d02ZRYLtxI4QriCw1cz2ux1');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('NYUh5d02ZRYLtxI4QriCw1cz2ux1', 7, '2022-05-29', 'NYUh5d02ZRYLtxI4QriCw1cz2ux1');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('NYUh5d02ZRYLtxI4QriCw1cz2ux1', 8, '2022-05-30', 'NYUh5d02ZRYLtxI4QriCw1cz2ux1');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('NYUh5d02ZRYLtxI4QriCw1cz2ux1', 9, '2022-05-31', 'NYUh5d02ZRYLtxI4QriCw1cz2ux1');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('IYadf5AYZYZByyTTl1f5QqxOGx13', 15, '2022-05-27', 'IYadf5AYZYZByyTTl1f5QqxOGx13');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('IYadf5AYZYZByyTTl1f5QqxOGx13', 16, '2022-05-28', 'IYadf5AYZYZByyTTl1f5QqxOGx13');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('IYadf5AYZYZByyTTl1f5QqxOGx13', 17, '2022-05-29', 'IYadf5AYZYZByyTTl1f5QqxOGx13');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('IYadf5AYZYZByyTTl1f5QqxOGx13', 18, '2022-05-30', 'IYadf5AYZYZByyTTl1f5QqxOGx13');
    INSERT INTO public.daily_limit
    (uid, task_limit, task_date, updated_by)
    VALUES('IYadf5AYZYZByyTTl1f5QqxOGx13', 19, '2022-05-31', 'IYadf5AYZYZByyTTl1f5QqxOGx13');

- user = tasker
  - 

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
  ### Testing
  #### Unit test
    cd $GOPATH/togo
    go test unit_test.go -v
  #### Integration test
    cd $GOPATH/togo
    go test integration_test.go -v
      
