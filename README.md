# TOGO

## Installation

### I. Third Apps
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

### II. Start
#### - **`${Build}`**
```sh
$ make
```

#### - **`${Test}`**
```sh
$ make test
```

#### - **`${Start Server}`**
```sh
$ make run
```

## How to call API
<details>
  <summary>Sign Up</summary>

  ### cURL
  ```sh
  curl --location --request POST "127.0.0.1:3000/api/v1/signup/" \
  --header 'Content-Type: application/json; charset=utf-8' \
  --data-raw '{
      "Password": "Password",
      "Username": "Username"
  }' | json_pp
  ```
</details>

<details>
  <summary>Login</summary>

  ### cURL
  ```sh
  curl --location --request POST "127.0.0.1:3000/api/v1/login/" \
  --header 'Content-Type: application/json; charset=utf-8' \
  --data-raw '{
      "Password": "Password",
      "Username": "Username"
  }' | json_pp
  ```
</details>

<details>
  <summary>Get List Tasks</summary>

  ### cURL
  ```sh
  curl --location --request GET "127.0.0.1:3000/api/v1/tasks/" \
  --header 'Content-Type: application/json; charset=utf-8' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJfaWQiOjF9LCJleHAiOjE2NDAwOTU5NDEsImlzcyI6IlNow6FuZW5vaSJ9.lQxr1wD3e4r1EVFmJ-Ir2nr90qC-2H3pR2iwxzd7PvU' | json_pp
  ```
</details>

<details>
  <summary>Create Task</summary>

  ### cURL
  > Body support various ways with this format `String: String`
  ```sh
  curl --location --request POST "127.0.0.1:3000/api/v1/tasks/" \
  --header 'Content-Type: application/json; charset=utf-8' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJfaWQiOjF9LCJleHAiOjE2NDAwOTU5NDEsImlzcyI6IlNow6FuZW5vaSJ9.lQxr1wD3e4r1EVFmJ-Ir2nr90qC-2H3pR2iwxzd7PvU' \
  --data-raw '{
    "assignments": "@shanenoi",
    "tasks": "prepair for a interview in Manabie"
  }' | json_pp
  ```
</details>

### If you are not *familiar* with `Console UI`
> Open Postman and click `ctrl + O`, then click `raw test` and paste above cURL, Postman will create a request base on this cURL

## Architecture
> I used to apply this Architecture for various Business Use Case, this one make you are more comfortable when add new logic or apply other database to specify services, this prefer for implementing gRPC dua with HTTP in our services

![](codebase.svg)
