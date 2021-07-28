#!/bin/bash
echo "Run database migrate"
goose -dir=/internal/database/migrations/ postgres "host=$DB_HOST user=$DB_USERNAME dbname=$DB_NAME password='$DB_PASSWORD' sslmode=$DB_SSL_MODE" up 
/go/bin/$APP