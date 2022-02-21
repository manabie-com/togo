echo $PWD
cd $PWD/internal/

go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out
