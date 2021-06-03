FROM postgres:11.6

COPY docker/db/postgresql.conf /etc/postgresql/postgresql.conf
RUN mkdir -p /var/lib/postgresql-static/data
ENV PGDATA /var/lib/postgresql-static/data
CMD ["-c", "config_file=/etc/postgresql/postgresql.conf"]
