### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- Install [docker](https://docs.docker.com/engine/install/ubuntu/)
- Install [docker-compose](https://docs.docker.com/compose/install/)
- Setup [make](https://www.gnu.org/software/make/) (optional)
- Run `make run` or `docker-compose up -d --build` if you don't have `make`
- Import Postman collection from docs to check example (at ./docs directory)

What I use:
- Using Golang, docker, make, [mockgen](https://github.com/golang/mock) to make this project better

What I did:
- [Setup docker for auto build project, run with postgres, redis](https://github.com/surw/togo/pull/3/commits/ec01d2db7047c5224cca3277072f4be44795722e)
- Refactoring: 
  - [Splitting out service, router (exporter), database.](https://github.com/surw/togo/pull/3/commits/ec01d2db7047c5224cca3277072f4be44795722e)
  - [Loose coupling component.](https://github.com/surw/togo/pull/7/commits/2ade71cf4ab50a48c602111b2522843a2221f015)
- [Switch database to postgres](https://github.com/surw/togo/pull/5/commits/a068839b8fdd819c0c29871b5dbbfd0ed0959749)
- [Using redis to implement limit control](https://github.com/surw/togo/pull/6/commits/c61539c4ce5f66536afdd6485603137ac8e1c952)
- [Add simple mock test](https://github.com/surw/togo/pull/8/commits/012e26694bdaefd987d7ef78ba5de54257289b7a)

What I think:
- This project quite generic, so I don't make any business change.
- There are a lot of places which not using proper type (for e.g: sql.NullString).
- I'm only implemented one testcase to show that I know how to write a test (as the requirement).
- Lastly, I wanna to say thank you who read to here :)
