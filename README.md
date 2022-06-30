### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### Folder structure

```bash
.
├── cmd
│   └── server
│       └── main.go
├── d.md
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── delivery
│   │   ├── auth.go
│   │   ├── auth_test.go
│   │   ├── handler.go
│   │   ├── plan.go
│   │   ├── plan_test.go
│   │   ├── task.go
│   │   ├── task_test.go
│   │   ├── user.go
│   │   └── user_test.go
│   ├── domain
│   │   └── domain.go
│   ├── entities
│   │   ├── task.go
│   │   └── user.go
│   ├── integration
│   │   └── integration_test.go
│   ├── middleware
│   │   └── middleware.go
│   ├── repository
│   │   ├── db.go
│   │   ├── task.go
│   │   ├── task_test.go
│   │   ├── user.go
│   │   └── user_test.go
│   ├── routes
│   │   ├── auth.go
│   │   ├── plan.go
│   │   ├── routes.go
│   │   ├── task.go
│   │   └── user.go
│   └── usecase
│       ├── task.go
│       ├── task_test.go
│       ├── user.go
│       └── user_test.go
├── LICENSE
├── Makefile
├── migrations
│   ├── create_tasks_table.sql
│   └── create_users_table.sql
├── pkg
│   ├── cryto.go
│   ├── json.go
│   ├── jwt.go
│   ├── mock
│   │   └── mock.go
│   └── utils.go
└── README.md
```