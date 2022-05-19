## Contents
- [Usage](#usage)
   - [Run normally](#run-normally)
   - [Run through docker](#run-through-docker)
- [Call API](#call-api)
   - [With curl command](#curl-command)
   - [With swagger](#swagger)
- [Tests locally](#tests-locally)
   - [Check code format](#check-code-format)
   - [Run unit test only](#run-unit-test-only)
   - [Run integration test only](#run-integration-test-only)
   - [Run all test](#run-all-test)
   - [Run coverage test](#run-coverage-test)

## Usage

- You must create file `.env` for the `first time` (in this case, I will put file.env into git).
- Description variable of env file:

| Environment    | Is_Required | Description                                                                                           |
| -------------- | ----------- | ----------------------------------------------------------------------------------------------------- |
| DATABASE_URL   | `true`      | MongoDB connection string with the standard URI connection scheme has the form: mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb] |
| PORT           | `true`      | Project's port                                                                                        |


- There is several way to run the project:

   ### Run normally:
   - You must use `file .env` and `mongoose` are running local with port `27017` and `Nodejs must be > 16.x`:
      ```env
      DATABASE_URL=mongodb://localhost:27017
      PORT=3000
      ```
   - Then run `cmd`:

      ```bash
      yarn
      yarn start
      ```
   - If you don't have `yarn` or you prefer to run with `npm`:
      ```bash
      npm i
      npm start
      ```

   ### Run through docker:
   - You must use `file .env` 
      ```env
      DATABASE_URL=mongodb://docker:27017
      PORT=3000
      ```
   - Then run `cmd`

      ```bash
      sudo docker-compose up
      ```


## Call API
- I have several choice for you:

   ### `curl` command: 
   - Check health of server: 
      ```bash
      curl -v http://localhost:3000/api/ping  
      ```

   ### Swagger:
   - Go to http://localhost:3000/api/docs

## Tests locally
- Before doing any test please install the package first with command:
   ```bash
   yarn
   ```
- Or with `npm`:
   ```bash
   npm i
   ```
   
- Then I will give you some choice again:

   ### Check code format:
   ```bash
   yarn lint
   ```
   - Or with `npm`:
   ```bash
   npm run lint
   ```
   
   ### Run unit test only:
    ```bash
   yarn test:unit
   ```
   - Or with `npm`:
   ```bash
   npm run test:unit
   ```

   ### Run integration test only:
   ```bash
   yarn test:integration
   ```
   - Or with `npm`:
   ```bash
   npm run test:integration
   ```

   ### Run all test:
   ```bash
   yarn test
   ```
   - Or with `npm`:
   ```bash
   npm run test
   ```

   ### Run coverage test:
   ```bash
   yarn test:coverage
   ```
   - Or with `npm`:
   ```bash
   npm run test:coverage
   ```