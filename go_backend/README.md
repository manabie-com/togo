### Dependency
- Nodemon
- postgresql

### Setup psql
- Setup psql account
  - CREATE ROLE testuser CREATEDB WITH LOGIN PASSWORD 'P@55w0rd'; 
  - CREATE DATABASE testdb;
- Run “create_table.sql” script included in folder

### How to run locally
- go to “go_backend” directory
- run via “make all”
  - for running test cases: “go test main_test.go -v”

- I made a simple approach to complete the exam

