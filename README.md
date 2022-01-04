# Production-ready Simple Todo Service Demo

![ci-test](https://github.com/lavantien/togo/actions/workflows/ci.yml/badge.svg?branch=master)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-78%25-brightgreen.svg?longCache=true&style=flat)</a>

## Requirements

1. [X] Implement core business logic
2. [X] User sign up/login with authentication/authorization via Paseto bearer token
3. [X] Task endpoints: create (with transaction, rate limitted), view lists (user bound)
4. [X] Tests cover all critical points
5. [X] GitHub Actions CI
6. [X] Docker Compose single point bootstraping

### Dissect Business Logics

#### Task Creation

<details>
	<summary>See details</summary>

```txt
a user?
	username
	dailyCap
	dailyQuantity

	> createUser
	> viewUsers
	> loginUser
	> updateUserDailyQuantity

	> createTask
	> listTasks
	(> editTaskByName)
	(> deleteTaskByName)
	> countTasksCreatedToday

	> updateUserDailyCap (admin only)

a task?
	id
	name - unique per owner
	content
	owner
	quantity

problem?
	a user create some tasks, the number of tasks must be within the user's dailyCap

logic?
	which user?
	provide task name
	provide task content
	how to check for new day to reset count?
	check if dailyQuantity+1 > dailyCap?
			return 500 ISE - "daily limit exceed"
		else
			if dailyQuantity != countTasksCreatedToday()?
				dailyQuantity=0
			dailyQuantity++
			updateUserDailyQuantity()
			return 200 OK
	fraud protection?
		tasks are not really deleted, they only marked as deleted but retains creation day, so we can count based on that day

params:
	username
	name
	content

result:
	user
	task
```

</details>

### API Walkthrough

<details>
	<summary>See details</summary>

![API](/resource/readme/api.png "API")

- Create a user:

```bash
curl http://localhost:8080/users -H "Content-Type: application/json" -d '{"username":"tienla","full_name":"Tien La","email":"tien@email.com","password":"matkhau"}' | jq
# Result
{
  "username": "tienla",
  "full_name": "Tien La",
  "email": "tien@email.com",
  "daily_cap": 0,
  "daily_quantity": 0,
  "password_change_at": "0001-01-01T00:00:00Z",
  "created_at": "2022-01-03T23:22:09.197164Z"
}
```

- Login as admin:

```bash
curl http://localhost:8080/users/login -H "Content-Type: application/json" -d '{"username":"admin","password":"secret"}' | jq
# Result
{
  "access_token": "v2.local.nvAS-aI1sdsEnexIp3K77Qo0jtXFb_XS9cQUCeYgcAEJzY3nwG97chAFfbsCMygdvxR_Ube3Hp_6kbR96EFrWc1PRu1yunkRvEnhTrjhtmr0Vur-kX_oaFIeMqFYGwz8cHCgT2oX53_PZi_I7_N27iudNA6jE3wiwTpokFd0euaOSefxaAzFYpAwu94bB-30msBqiDgTR6ouDzB42dC1jhMp3rdRsOHDLV_xeiSBHt5UKqtA_aYp51G8dzMTTSdPXqZ0DAc3lNIn5q-T8g.bnVsbA",
  "user": {
    "username": "admin",
    "full_name": "Admin",
    "email": "admin@email.com",
    "daily_cap": 10,
    "daily_quantity": 0,
    "password_change_at": "0001-01-01T07:00:00Z",
    "created_at": "2021-12-26T22:22:49.644Z"
  }
}

# Save token for later use
ADMIN_TOKEN='v2.local.nvAS-aI1sdsEnexIp3K77Qo0jtXFb_XS9cQUCeYgcAEJzY3nwG97chAFfbsCMygdvxR_Ube3Hp_6kbR96EFrWc1PRu1yunkRvEnhTrjhtmr0Vur-kX_oaFIeMqFYGwz8cHCgT2oX53_PZi_I7_N27iudNA6jE3wiwTpokFd0euaOSefxaAzFYpAwu94bB-30msBqiDgTR6ouDzB42dC1jhMp3rdRsOHDLV_xeiSBHt5UKqtA_aYp51G8dzMTTSdPXqZ0DAc3lNIn5q-T8g.bnVsbA'
```

- Update the daily cap of the newly created user to 2:

```bash
curl http://localhost:8080/admin/setDailyCap -H "Authorization: Bearer $ADMIN_TOKEN" -H "Content-Type: application/json" -d '{"username":"tienla","daily_cap":2}' | jq
# Result
{
  "username": "tienla",
  "full_name": "Tien La",
  "email": "tien@email.com",
  "daily_cap": 2,
  "daily_quantity": 0,
  "password_change_at": "0001-01-01T00:00:00Z",
  "created_at": "2022-01-03T23:22:09.197164Z"
}
```

- Now login as the newly create user:

```bash
curl http://localhost:8080/users/login -H "Content-Type: application/json" -d '{"username":"tienla","password":"matkhau"}' | jq
# Result
{
  "access_token": "v2.local.K5YFZ2P-9cZoN_HtS8olPqHZ88ku1whq3cUS3R88y6qlJJlACtBwhcpEQrW8fgZVxJQlxai-p68gv9vIJuP8K0f111uBh6Mceq2bhP9T64_oS4SbPzKIgXlG_pase7H-QYytWFzdo3rjCRY19GW-Ev3c-NgHlcy9GGlECgrtKU053JxVv54GFUpkovL8oXvtk-BYUOHywJH-na3126GyqT1G9h-EwYMrgV_PrcFfWOVGdSwhzN568Ta9rClUOKlvdTlTikaH-5laih1YKbQ.bnVsbA",
  "user": {
    "username": "tienla",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "daily_cap": 2,
    "daily_quantity": 0,
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:22:09.197164Z"
  }
}

# Save the token for later use
TOKEN='v2.local.K5YFZ2P-9cZoN_HtS8olPqHZ88ku1whq3cUS3R88y6qlJJlACtBwhcpEQrW8fgZVxJQlxai-p68gv9vIJuP8K0f111uBh6Mceq2bhP9T64_oS4SbPzKIgXlG_pase7H-QYytWFzdo3rjCRY19GW-Ev3c-NgHlcy9GGlECgrtKU053JxVv54GFUpkovL8oXvtk-BYUOHywJH-na3126GyqT1G9h-EwYMrgV_PrcFfWOVGdSwhzN568Ta9rClUOKlvdTlTikaH-5laih1YKbQ.bnVsbA'
```

- View user's information (only admin can see all):

```bash
curl http://localhost:8080/users?page_id=1&page_size=5 -H "Authorization: Bearer $TOKEN" | jq
# Result
[
  {
    "username": "tienla",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "daily_cap": 2,
    "daily_quantity": 0,
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:22:09.197164Z"
  }
]
```

- Create a task:

```bash
curl http://localhost:8080/tasks -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"name":"task 1","content":"This is task number 1"}' | jq
# Result
{
  "user": {
    "username": "tienla",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "daily_cap": 2,
    "daily_quantity": 1,
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:22:09.197164Z"
  },
  "task": {
    "id": 8,
    "name": "task 1",
    "owner": "tienla",
    "content": "This is task number 1",
    "deleted": false,
    "content_change_at": "0001-01-01T00:00:00Z",
    "deleted_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:24:21.546377Z"
  }
}
```

- Create another task:

```bash
curl http://localhost:8080/tasks -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"name":"task 2","content":"This is task number 2"}' | jq
# Result
{
  "user": {
    "username": "tienla",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "daily_cap": 2,
    "daily_quantity": 2,
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:22:09.197164Z"
  },
  "task": {
    "id": 9,
    "name": "task 2",
    "owner": "tienla",
    "content": "This is task number 2",
    "deleted": false,
    "content_change_at": "0001-01-01T00:00:00Z",
    "deleted_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:24:46.500126Z"
  }
}
```

- Create yet another task (this time it's over the daily limit):

```bash
curl http://localhost:8080/tasks -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{"name":"task 3","content":"This is task number 3"}' | jq
# Result
{
  "error": "daily limit exceed"
}
```

- View all the tasks that you've successfully created (only 2 tasks):

```bash
curl http://localhost:8080/tasks?page_id=1&page_size=5 -H "Authorization: Bearer $TOKEN" | jq
# Result
[
  {
    "id": 8,
    "name": "task 1",
    "owner": "tienla",
    "content": "This is task number 1",
    "deleted": false,
    "content_change_at": "0001-01-01T00:00:00Z",
    "deleted_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:24:21.546377Z"
  },
  {
    "id": 9,
    "name": "task 2",
    "owner": "tienla",
    "content": "This is task number 2",
    "deleted": false,
    "content_change_at": "0001-01-01T00:00:00Z",
    "deleted_at": "0001-01-01T00:00:00Z",
    "created_at": "2022-01-03T23:24:46.500126Z"
  }
]
```

</details>

### Project Structure

<details>
	<summary>See details</summary>

```bash
.
├── api
│   ├── admin.go
│   ├── admin_test.go
│   ├── main_test.go
│   ├── middleware.go
│   ├── middleware_test.go
│   ├── server.go
│   ├── task.go
│   ├── task_test.go
│   ├── user.go
│   ├── user_test.go
│   └── validator.go
├── common
│   └── model
│       └── user_response.go
├── db
│   ├── migration
│   │   ├── 000001_init_schema.down.sql
│   │   └── 000001_init_schema.up.sql
│   ├── mock
│   │   └── store.go
│   ├── query
│   │   ├── task.sql
│   │   └── user.sql
│   └── sqlc
│       ├── db.go
│       ├── main_test.go
│       ├── models.go
│       ├── querier.go
│       ├── store.go
│       ├── store_test.go
│       ├── task.sql.go
│       ├── user.sql.go
│       └── user_test.go
├── .github
│   └── workflows
│       └── ci.yml
├── log
│   └── createTaskTx_decision_log.txt
├── resource
│   ├── readme
│   │   ├── api.png
│   │   └── togo.png
│   ├── debug.pgsql
│   └── togo.sql
├── token
│   ├── maker.go
│   ├── paseto_maker.go
│   ├── paseto_maker_test.go
│   └── payload.go
├── util
│   ├── config.go
│   ├── config_test.go
│   ├── fullname.go
│   ├── fullname_test.go
│   ├── password.go
│   ├── password_test.go
│   ├── random.go
│   └── random_test.go
├── app.env
├── coverage_badge.png
├── coverage.md
├── coverage.out
├── docker-compose.yml
├── Dockerfile
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── profile.cov
├── README.md
├── sqlc.yaml
├── start.sh
├── tree.txt
└── wait-for.sh

15 directories, 62 files
```

</details>

### Database UML

![Database UML](/resource/readme/togo.png "Database UML")

## Technology Stack

- **Go 1.17**: *Leverage the standard libraries as much as possible*
- **SQLc**: *Generates efficient native SQL CRUD code*
- **PostgreSQL**: *RDBMS of choice because of faster read due to its indexing model and safer transaction with better isolation levels handling*
- **Golang-Migrate**: *Efficient schema generating, up/down migrating*
- **Viper**: *Add robustness to configurations*
- **Gin**: *Fast and have respect for native net/http API*
- **GoMock**: *Generates mocks of about anything*
- **Paseto Token**: *Better choice than JWT because of enforcing better cryptographic standards and debloated of useless information*
- **Docker** + **Docker-Compose**: *Containerization, what else to say ...*
- **Github Actions CI**: *Make sure we don't push trash code into the codebase*
- **GopherBadger**: *That nice coverage badge in this README*

## Philosophy and Architecture

- **Adaptive Minimalism**: *I always keep it as simple as possible, but with a highly decoupled structure it's ensure high adaptivity and extensibility, on top of that minimal solid head start. Things are implement only when they're absolutely needed*

## Get Start

### Manual Way

- Spin up a postgres container: `make postgres`
- Create database: `make createdb`
- Database migration via Golang-Migrate: `make migrateup`
- Generate Go DB layer via SQLc: `make sqlc`
- Generate Mocks layer via GoMock: `make mock`
- Run tests and populate data: `make test`
- Run the backend: `make server`
- See code coverage report: `make cover`
- Make cover badger: `make badger`
- Continue with the **API Walkthrough**
- Check the `Makefile` for more commands!

### Or With Docker Compose

- Spin up the containers: `docker-compose up`
- Run tests and populate data: `make test`
- Continue with the **API Walkthrough**
- Shut down the containers: `docker-compose down`

## Struggle with Installation? Here is my detailed guide

<details>
	<summary>More details</summary>

- [**Golang**](https://go.dev/doc/install):

```bash
# Go to go.dev/dl and download a binary, in this example it's version 1.17.5

sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz

# Add these below to your .bashrc or .zshrc
export GOPATH=/home/<username>/go
export GOBIN=/home/<username>/go/bin
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOBIN
```

- [**Docker**](https://docs.docker.com/engine/install/ubuntu/):

```bash
sudo apt remove docker docker-engine docker.io containerd runc

sudo apt update

sudo apt install apt-transport-https ca-certificates curl gnupg lsb-release software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update

apt-cache policy docker-ce

sudo apt install docker-ce docker-ce-cli containerd.io

sudo usermod -aG docker $USER

newgrp docker

# Restart the machine then test the installation

docker run hello-world

# On older system you also need to activate the services

sudo systemctl enable docker.service

sudo systemctl enable containerd.service
```

- [**Docker-Compose**](https://docs.docker.com/compose/install/):

```bash
# Check their github repo for latest version number
sudo curl -L "https://github.com/docker/compose/releases/download/v2.2.2/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose

# To self-update docker-compose
docker-compose migrate-to-labels
```

- [**Golang-Migrate**](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate):

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- [**SQLc**](https://docs.sqlc.dev/en/latest/overview/install.html):

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

- [**GoMock**](https://github.com/golang/mock):

```bash
go install github.com/golang/mock/mockgen@latest
go get github.com/golang/mock/mockgen
```

- [**Viper**](https://github.com/spf13/viper):

```bash
go install https://github.com/spf13/viper@latest
```

- [**Gin**](https://github.com/gin-gonic/gin#installation):

```bash
go install github.com/gin-gonic/gin@latest

go get -u github.com/gin-gonic/gin
```

- [**Paseto**](https://github.com/o1egl/paseto):

```bash
go get -u github.com/o1egl/paseto
```

- [**JWT**](https://github.com/golang-jwt/jwt):

```bash
go get -u https://github.com/golang-jwt/jwt
```

- [**GopherBadger**](https://github.com/jpoles1/gopherbadger):

```bash
go install github.com/jpoles1/gopherbadger@latest
```

- [**CURL**](https://curl.se/download.html) + [**JQ**](https://stedolan.github.io/jq/) + [**Chocolatery**](https://docs.chocolatey.org/en-us/choco/setup) + [**Make**](https://community.chocolatey.org/packages/make):

```bash
sudo apt install curl jq

# These tools are needed only for Windows users

# Run this in an Admin cmd to install Chocolatery first
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

# Then install GNU-Make, cURL, and jq via Chocolatery in Admin pwsh
choco install make curl jq
```

</details>
