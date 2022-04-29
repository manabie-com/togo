# togo

- this is service for node writing by typescript

## Version specification

- node: v12.22.0
- TypeScript v4.5.5

## IDE recommendations and setup

- VSCode IDE
- Eslint vscode plugin
- Prettier vscode plugin for linting warnings (auto fix on save)
- Add the following setting in vscode settings.json

```json
"eslint.autoFixOnSave": true
```

## Dev setup

- Install all the dependencies using `npm install`
- To run the server with watch use `npm run start:dev`

## Test

- Unit Test: We are using Jest for assertion and mocking
- To run the test cases use `npm run test`
- To get the test coverage use `npm run test:cov`

## Git Hooks

The seed uses `husky` to enable commit hook.

### Pre commit

Whenever there is a commit, there will be check on lint, on failure commit fails.

## ENV variables

- create .env file for all config.

```none
SERVICE_NAME=seeds-node
HOST=localhost
PORT=3000
LOG_LEVEL=info
```

## Misc

- Swagger API is at <http://localhost:3000/documentation>
