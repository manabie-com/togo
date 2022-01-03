# Production-ready Simple Todo Service Demo

![ci-test](https://github.com/lavantien/togo/actions/workflows/ci.yml/badge.svg?branch=master)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-79%25-brightgreen.svg?longCache=true&style=flat)</a>

## Requirements

1. [X] Implement core business logic
2. [X] User sign up/login with authentication/authorization via Paseto bearer token
3. [ ] Task endpoints: create (with transaction, rate limitted), view lists (user bound)
4. [ ] Tests cover all critical points
5. [X] GitHub Actions CI
6. [ ] Docker Compose bootstrap

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
			return 403 Fobbiden - "daily limit exceed"
		else
			if dailyQuantity != countTasksCreatedToday()?
				dailyQuantity=0
			dailyQuantity++
			updateUserDailyQuantity()
			return 201 Created
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

- **Adaptive Minimalism**: *I always keep it as simple as possible, but with a highly decoupled structure we ensure high adaptivity and extensibility, on top of that minimal solid head start. Things are implement only when they're absolutely needed*

## Get Start

- Spin up a postgres container: `make postgres`
- Create database: `make createdb`
- Database migration via Golang-Migrate: `make migrateup`
- Generate Go DB layer via SQLc: `make sqlc`
- Generate Mocks layer via GoMock: `make mock`
- Run tests!: `make test`
- Run the backend: `make server`
- See code coverage report: `make cover`
- Check the `Makefile` for more commands!

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
