# **Tasks to do**:
    - [x] Make plan.
    - [x] Setup Dev environment: Nginx, React App, Golang Server, Redis Cache, Postgres Database.
    - [x] Write Integration Test.
    - [x] Structure Server Code. 
    - [x] Create Log Service
    - [x] Write Unit Test.
    - [x] Write Logic Code for basic features
    - [x] Write Logic Code for main requirements
    - [x] Structure React Code with Redux
    - [ ] UI for List/Create todo feature
    - [ ] Write Unit Test for List/Create todo feature
    - [ ] UI for Login feature
    - [ ] Write Unit Test for Login feature
    - [ ] UI for Logic feature

# How to run 
    - Please type `make run` and wait for the building process. The Server API will be available at localhost:8080/api. The Web App will be on localhost:8080.

# What I Have Done:
    - Limit Create Task For 5 times per 24 hour, safe for concurrent access and not reduce latency per request so much.
    - Write Integration Test
    - Structure code by implementing clean code architecture.
    - Split `services` layer to `use case` and `transport` layer, Make logic code happen on handler, usecase, storage layer, reduce boilerplate code of transport layer.
    - Write unit test for `services` layer
    - Change from using `SQLite` to `Postgres` with `docker-compose`
    - Write unit test for `storages` layer
    - Redesign database
    - Create structure for frontend with redux  
# What I Miss:
    - Frontend app not work.
# What I Want to improve
    - Write script to gen transport layer
