## Create PostgreSQL server using Docker

```batch
docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=fUaECg4SIei7 -v postgres:/var/lib/postgresql/data postgres:14
```