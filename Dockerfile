

FROM postgres

# config env variables to dockerize
ENV POSTGRES_USER your_pg_user
ENV POSTGRES_PASSWORD your_pg_password 
ENV POSTGRES_DB your_db_name

COPY migrations/migrate.sql /docker-entrypoint-initdb.d/
VOLUME ./postgres-data:/var/lib/postgresql/data