#export DB_URL="postgres://postgres:abcd1234@golangvietnam.com:54321/rf_sms?sslmode=disable"
echo $PWD
cd $PWD/internal/

go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out
