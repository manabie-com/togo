#!/bin/bash

set -e
set -u

echo "Creating user and database '$POSTGRES_DB_TESTING' with user '$POSTGRES_USER' in 2 s"

sleep 2

psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" <<-EOSQL
	    CREATE DATABASE $POSTGRES_DB_TESTING;
	    GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB_TESTING TO $POSTGRES_USER;
EOSQL

echo "Database '$POSTGRES_DB_TESTING' created"
