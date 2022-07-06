FROM postgres
ENV POSTGRES_PASSWORD manabie
ENV POSTGRES_DB togo
COPY script.sql /docker-entrypoint-initdb.d/
VOLUME ./postgres-data:/var/lib/postgresql/data