### Init

Default port for mongodb and api is 27017 and 81 respectively

> cp .env.example .env \
> docker compose up \
> docker compose exec api go run seeder/main.go

### Seeding

> docker compose exec api go run seeder/main.go

### Invoking

```
curl -X POST \
-d '{"title": "Test Task", "assignee_email": "ptrung@manabie.test"}' \
-k http://localhost:81/task
```

### Missing

I could not spend more time to complete testing implement but we can exchange about this one in technical interview.
