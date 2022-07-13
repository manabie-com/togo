### Init

> cp .env.example .env
> docker compose up
> docker compose exec api go run seeder/main.go

### Seeding

> docker compose exec api go run seeder/main.go

### Invoking

```
curl -X POST \
-d '{"title": "Test Task", "assignee_email": "ptrung@manabie.test"}' \
-k http://localhost:81/task
```
