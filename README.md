# Minimal Todo Service with Golang
- CRUD for Todo Task
- Support JWT
- User can create todo task with quota limit per day.
- There are two user roles in system: `admin` and `user`.
- Admin is only able to update quota setting for user.
  
# Install and run in local
1. `cp .env.example .env`
2. Update `.env` for database info.
3. Run `air`

# To run test locally
- go test ./tests/functional/ -v

# Endpoints
Refer `./docs/endpoint.txt`
