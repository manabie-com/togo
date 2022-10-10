FROM mysql:8.0.23

COPY ./database/docker/dev/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./database/docker/dev/migration/*.sql /docker-entrypoint-initdb.d/