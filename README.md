### Test
- You need docker and docker-compose installed on your machine
- Run `./test.sh`
- If success, the test reports will be put in tests/report folder. You can check html files for code coverage

### Run
- Run `./devup.sh` to start the application
- Configurations are stored at cmd/app/env folder
- The app is deployed with Prometheus for metric gathering, Grafana for metric dashboard and Postgresql for database
- Incase you need a username/password to log in, get into the container of the app and use `./cli add-admin username password`

### TODO
- Consider distributed cache for task adding with per day limit, SQL transaction is hard to write and optimize in this use case
- Restructure database model, add more indexes for sorting  
- Add logging