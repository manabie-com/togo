# Enhancement

## Implement business logic

Add limit n task per day.

## Refine project structure

We will use project structure from https://github.com/golang-standards/project-layout

## Refine Authorization

Follow [Syntax](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization)

`Authorization: <type> <credentials>`

We need add type `bearer` to this.

## Replace authorization handler

Replace with authorization interceptor

## Don't hardcode config parameter (jwt key....)

We will move config parameter (jwt key) to file config or env depend on which environment deploy app

## Encrypt password

Save with encrypted password

## Refactor table users

* We will add `username` to this table and then we will login with `username` and `password`. The `id` will generated
  random when user created .
* Replace created_date from string to timestamp

## Refactor api login

* Sensitive data should not be sent via a HTTP "GET", should always be sent via POST.
* Request data should be sent via body.

## Add new way return error

Currently, we return error throw http body

````go
json.NewEncoder(resp).Encode(map[string]string{
"error": err.Error(),
})
````

I refer http header rather than http body

## Add ORM

We will add `ent` or `gorm`, so we can use other database.

## Auto retry database connection

## Apply DI

apply dependency injection

## Use other http server

I don't like `switch case` when handle http request. I think, we will add `chi` or `gin`...

## Add log

We will add OpenTracing

## Add monitoring

## Add health check

## Graceful shutdown

## Handle request/response with converter

We will encode/decode base on http header (`Content-Type`, `Accept`)

## Support Etag

## Add API docs

We will use swagger 
