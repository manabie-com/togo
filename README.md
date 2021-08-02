# Togo setup instruction

## Installing dependencies

Run this command in the terminal then run the script:
```bash
$ chmod +x install_dependencies.sh
./install_dependencies.sh
```

## Setting up Docker for the database

Setting up docker: `https://docs.docker.com/engine/install/ubuntu/`

Start Docker, then run this command in the terminal:
```bash
$ docker-compose up -d
```
## Setting up the database

Database login credentials can be found in the `docker-compose.yml` file.

Login to database using pgadmin by accessing to `http://localhost:16543/` using credentials from `docker-compose.yml`:
```
Email address: admin@admin.com
Password: admin
```

Create a new server in pgadmin with the name `togo-db-local`.

After creating, right-click the server, select `Properties` and switch to the `Connetion` tab.

Update the connection with the data below:
```
Host name/address: togo-db-postgres
Port: 5432
Username: username
Password: password
```

### Starting up the server for the first time: 
In `main.go`, change `database.SyncDB(false)` to `database.SyncDB(true)`.

Explanation: This line of code will force the server to truncate & recreate tables based on the `models` files.

Start the server by running the `start_app_local.sh` script:
```bash
$ chmod +x install_dependencies.sh
./start_app_local.sh
```
Once the database is migrated,close the server, revert the code change, then start the server again.

### Starting up the server: 
Start the server by running the `start_app_local.sh` script:
```bash
./start_app.local.sh
```