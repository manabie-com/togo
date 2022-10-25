FROM mysql:8.0.23

COPY ./migration/*.sql  /docker-entrypoint-initdb.d/