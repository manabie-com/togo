# TODO API

Tiny api to create todo list. Made with Ruby on Rails.


## Main idea

```ruby
def count_request
  redis.set(client_id, 0, ex: 1.day) if redis.get(client_id).nil?
  redis.incr(client_id)
end

def client_id
  "#{client_remote_ip}:#{date}"
end
```
- For each request, find (or create anew) user based on the IP address. This is the simplest method for tracking user as there is no need (and requirement) for creating users.
- Track the request each time it reaches the server on Redis.
- The key for each request counter is based on user ip address and the current date, so it is reset each day.

## Set-up and Installation (macos only)

Dependency: redis, ruby 3.0.2, rails 6 (suggested installation method: `brew` & `asdf`)

```
git clone https://github.com/tuang3142/togo.git
cd togo
bin/setup
bin/rails s
```
Optional: change the value in `config/local_env.yml` for updating post request daily limit.

## Example API calls

```
❯ curl -d title='title' 'http://localhost:3000/api/v1/todos'
{"status":201,"message":"Todo created","data":{"id":35,"title":"title","content":null,"user_id":12,"created_at":"2022-04-26T13:30:37.521Z","updated_at":"2022-04-26T13:30:37.521Z"}}

❯ curl -d title='shopping list' -d content='apple,milk,icecream' 'http://localhost:3000/api/v1/todos'
{"status":201,"message":"Todo created","data":{"id":36,"title":"shopping list","content":"apple,milk,icecream","user_id":12,"created_at":"2022-04-26T13:31:38.240Z","updated_at":"2022-04-26T13:31:38.240Z"}}%

❯ curl -d title='study' -d 'http://localhost:3000/api/v1/todos'
{"status":429,"message":"Too many requests"}
```

## Running test

```
❯ bundle exec rspec
....

Finished in 0.1784 seconds (files took 1.32 seconds to load)
4 examples, 0 failures
```

The main test logic can be found in `spec/requests/todo_request_spec.rb`