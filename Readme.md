### Infra

- create network 
 ``` bash
    make network
 ```
- Start postgres in docker
 ```bash
    make postgres
 ```
- Create database

 ```bash
    make createdb
 ```

- db migration up

```bash
  make migrateup  
```
- db migration down
```bash
  make migratedown
```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```