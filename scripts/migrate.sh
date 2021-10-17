export POSTGRESQL_URL='postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable'
./migrate -database ${POSTGRESQL_URL} -path db/migrations $1 $2
