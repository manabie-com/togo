# Todo API

## Prerequisites
[![Ruby Style Guide](https://img.shields.io/badge/Ruby-3.1.2-red)](https://www.ruby-lang.org/en/news/2022/04/12/ruby-3-1-2-released)
[![Ruby Style Guide](https://img.shields.io/badge/Rails-7.0.3-brightgreen)](https://rubygems.org/gems/rails)

## Setup on WSL: Linux/Ubuntu

### First-time setup
1. Install Ruby
2. Install Rails
3. Install PostgreSQL.

### Running app on localhost
1. Change the host in database.yml to 'localhost'.
2. **Run** the following command.
```bash
$ cd togo
```
```bash
$ bundle install
```
```bash
$ rails db:create db:migrate db:seed
```
```bash
$ rails s
```
The server will run on **localhost:3000**.

### Setup on Docker
1. Change the host in database.yml to 'db'.
2. **Run** the following command.
```bash
$ docker-compose up -d
```
#### To migrate and seed:
```bash
$ docker-compose run web rake db:migrate db:seed
```

### Run Tests
```bash
$ rspec spec/requests/api
```

#### Remarks:
> I love my solution because I followed the Ruby on Rails standards and bes practices. <br/> <br/>
> Anti patterns was practiced in my solution by not having an Overweight Models and Controllers. <br/>
> I also used interactor for my design pattern. It is a simple, single-purpose object. Interactors are used to encapsulate your application's business logic. Each interactor represents one thing that your application does.  <br/>
