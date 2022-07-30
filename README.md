# Todo App
Simple Todo App API

## Prerequisites

- `Java 17+`
- `Docker`
- `Docker-compose`

## Get Started

### Setup environment

```shell
docker-compose up -d
```

### Run project

```shell
./mvnw spring-boot:run 
```

### Sample Curls

---
**NOTE:**
Created 1 user for testing purposes:
- id: 1, name: 'uuhnaut69', limit_config: 10
---

- Create new todo

```shell

```

- Full API documentation can visit [here]()


### Testing

For running Unit Tests:

```shell
 ./mvnw test 
```

For running Integration Tests:

```shell
 ./mvnw verify -Pintegration 
```


### Todo

- [ ] Authentication
- [ ] User limit configuration
- [ ] Web Interface