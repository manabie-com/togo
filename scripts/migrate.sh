#!/bin/bash

export POSTGRESQL_URL='postgres://togo_user:togo_password@db:5432/togo_db?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path db/migrations $1 $2