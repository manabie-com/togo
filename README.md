# Architecture
![togo_user_domain_sequence_diagram](./docs/images/todo_context_diagram.png)
The project's design is based on microservice architecture and communicate via Restful

The system have 6 component:
- Reverse proxy & Load balancer (`Nginx`)
- User Service: manage user account and all setting
- Task Service:
  + task counter: `rate limiter` is based on `leaky bucket algorithm` with no `weight`
  + manage task by  user
- User Database: is a `Postgres` instance for User Service
- Task Database: is a `Postgres` instance for Task Service
- Distributed Lock: is a `Redis` instance for caching & concurrent lock in distributed system

# Source structures
```sh
└── com
    └── manabie
              └──todo
                    ├── config : infrastructure config
                    ├── constant : All constants values that can be used in the services
                    ├── controller : It is the public face of the application layer. It routes incoming requests and returns responses.
                    ├── entity : Object was mapped with database
                    ├── exception: Common exception
                    ├── model: Object was mapped with request/response and business model
                    ├── repository: interact with infrastructure to get resources for service layer
                    ├── service: The domain layer is responsible for encapsulating complex business logic, or simple business logic that is reused by multiple Controller
```
# Software usage
- Spring Boot
- Spring Reactive Webflux
>Spring Framework uses Project Reactor as the base implementation of its reactive support, and also comes with a new web framework, Spring WebFlux, which supports the development of reactive, that is, non-blocking, HTTP clients and services.
- Docker
>Deploying Our Microservices Using Docker

# Running the microservices
1. Run `mvn clean package` to build the applications.
2. Run `docker-compose up -d` to create the docker image locally and start the applications.
# How to use
# Todo