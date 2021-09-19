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

dev:
	@go run main.go
