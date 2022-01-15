### Document
Please read here : https://docs.google.com/document/d/1uMVX5KGPF9Vjc9VqGP50PmkcWhI0vhlBf5rsDL4BIY4/

### API

- GET /userid/date/taskname
  - 200 = success.
- PUT /userid/date/taskname
  - Block. 201 = written and 204 = overwritten, anything else = probably not create.
- DELETE /userid/date/taskname
  - Lock. 204 = deleted, anything else = probably not deleted.

### Usage

```
# put "onyou" in key "userid/20210101/taskname" (will 403 if it already exists)
curl -v -L -X PUT -d onyou localhost:3000/userid/20210101/taskname

# get key "userid/20210101/taskname" (should be "taskcontent")
curl -v -L localhost:3000/userid/20210101/taskname

# delete key "userid/20210101/taskname"
curl -v -L -X DELETE localhost:3000/userid/20210101/taskname

# put file in key "file.txt"
curl -v -L -X PUT -T /path/to/local/file.txt localhost:3000/userid/20210101/taskname

# get file in key "file.txt"
curl -v -L -o /path/to/local/file.txt localhost:3000/userid/20210101/taskname
```

### Local starting
```
docker-compose -f docker-compose-local.yaml up --build
```
### Intergration testing
```
docker-compose -f docker-compose-local.yaml run test
```

### Benchmark testing
```
go run ./tools/thrasher.go
```

### Unit testing
```
go test src/*.go
```
