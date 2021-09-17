init-env:
	# run docker-compose
	@docker-compose -f docker-compose.yaml up -d
	# clear go.mod & go.sum
	@rm go.mod go.sum
	# install environment go
	@go mod init
	@go mod tidy
	# install dbmate
	@brew install dbmate

setup-db:
	@export DATABASE_URL="postgres://postgres:postgres@localhost:5434/togo_postgres?sslmode=disable"
	@dbmate -d ./db/init up
dev:
	@go run main.go
