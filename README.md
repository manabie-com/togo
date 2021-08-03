# Manabie Togo

[![Go|GoDoc](https://godoc.org/github.com/hellofresh/health-go?status.svg)](https://godoc.org/github.com/hellofresh/health-go)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

## Features

- Change method login to POST instead of GET.
- Add method register (POST /regisger).
- Split services layer to domain and transport layer.
- Write unit test for the domain layer.
- Write unit test for the repository/store layer.

## Project Structure

```sh
.
├── cmd/
│   └── srv/        
│   │   └── main.go           #  main func 
│   └── migrate                
│   |   └── main.go           #  main func for migrate
├── configs/
│   ├── config.go             # contain config variable
├── internal/ 
│   ├── transport/            # contain transport layer
│   │   └── ...
|   ├── domain/               # contain bussiness layer 
│   │   └── ...
│   └── storages/             # contain model and repository/store layer
│   │   └── postgres/         
│   │   |   └── ...
│   │   └── redis/         
│   │   |   └── ...
│   │   └── entities.go       #  contain models and entities
│   │   └── task.go           #  contain interface  of task store
│   │   └── user.go           #  contain interface of user store
├── docs/ 
├── common/ 
│   └── errors/
│   |   └── ... 
│   └── constants/
│   |    └── storages/  
├── test/                     # e2e test 
├── └── ... 
├── utils/ 
├── Dockerfile
├── docker-compose.yml
├── docker-compose.dev.yml
├── .gitignore                # specifies intentionally untracked files to ignore
├── .env
├── docker.env
├──.dockerignore
├── go.mod 
├── go.sum
```

## Installation

Make sure you have Go installed ([download](https://golang.org/dl/)). Version `1.16` or higher is required.

Install make for start the server.

For Linux:

```sh
$sudo apt install make
```

For Macos:

```sh
$brew install make
```

## Start server

First of all, you must copy .env.example to .env:

```sh
$cp .env.example .env
```

Start server with cmd/terminal:

```sh
$make docker-dev     # start docker with dev environment
$make migrate
$make start
```

Start server with docker:

```sh
$make docker-start
```

Your app should now be running on [localhost:5050](http://localhost:5050/).

## Unit test

```sh
$make unit-test
```

## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
