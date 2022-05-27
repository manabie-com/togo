
.PHONY: test
test:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: integrated_test
integrated_test:
	go test ./...