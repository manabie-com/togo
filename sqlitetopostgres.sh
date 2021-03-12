#!/bin/bash

SQLITE_DB_PATH='./data.db'
PG_DB_NAME=togo
PG_USER_NAME=phuonghau
PG_PORT=8899
PG_HOST=localhost

SQLITE_DUMP_FILE="sqlite_data.sql"

sqlite3 $SQLITE_DB_PATH .dump > $SQLITE_DUMP_FILE

# PRAGMAs are specific to SQLite3.
sed -i '/PRAGMA/d' $SQLITE_DUMP_FILE
# Convert sequences.
sed -i '/sqlite_sequence/d ; s/integer PRIMARY KEY AUTOINCREMENT/serial PRIMARY KEY/ig' $SQLITE_DUMP_FILE
# Convert column types.
sed -i 's/datetime/timestamp/g ; s/integer[(][^)]*[)]/integer/g ; s/text[(]\([^)]*\)[)]/varchar(\1)/g' $SQLITE_DUMP_FILE

createdb -U $PG_USER_NAME -h $PG_HOST -p $PG_PORT $PG_DB_NAME 
psql -U $PG_USER_NAME -h $PG_HOST -p $PG_PORT $PG_DB_NAME < $SQLITE_DUMP_FILE
