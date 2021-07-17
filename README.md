# Togo

## Startup
Simply

```shell
go get
go run main.go
```

## Run test
```shell
go test ./internal/...
```

With coverage:
```shell
go test -cover ./internal/...
```

## What I did
- I live and breathe TDD, so of course: Unit Testing
  - Feature/Integration Testing also
    - In my definition, Feature Test = HTTP Test, Integration Test = Test n Endpoints for a whole flow.
- Improved some parts of this project
- Password Encryption with BCRYPT
  - Security is a must
- Check Max Todo
- Create another DB for Unit Testing.

Finished this test in like 2h30m.

### Coverage result
```text
ok      github.com/manabie-com/togo/internal/services   8.542s  coverage: 85.7% of statements
ok      github.com/manabie-com/togo/internal/storages/sqlite    (cached)        coverage: 85.1% of statements
ok      github.com/manabie-com/togo/internal/helpers    (cached)        coverage: 100.0% of statements
```

## What did I miss
- Migrate SQLite to PostgreSQL.
  - Basically I feel like I don't have much knowledge about Go and if I do it, might cost a lot of time.

## What did I feel after the test
- I'm a grammar-nazi myself, and I don't feel comfortable to read the original README file.
  - Tips: you can use Grammarly to improve it.
- Overall, an ok test. 
  - It would be absolutely better if I can choose a specific language (the lang that I feel most comfortable to work with) to do the test.
  - Don't get the idea wrong, I know that you're hiring for a `Go` guy (but also opening for other languages). I'm willing to learn & working with Go as well.


## Final
Thanks, looking forward to the result.