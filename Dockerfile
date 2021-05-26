FROM postgres

# Use user/MyPassword1 as user/password credentials
ENV POSTGRES_PASSWORD=MyPassword1

WORKDIR /docker-entrypoint-initdb.d
ADD togo.sql /docker-entrypoint-initdb.d

EXPOSE 5432
