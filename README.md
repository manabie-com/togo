1. How to run?
  - Frontend Source
    - yarn install
    - yarn start:fullstack
    - Login api web:
      - Username: firstUser
      - Password: example

  - Backend Source:
    - npm install
    - npm start
    - Login api with:
      - Username: firstUser
      - Password: example
    
  - Unit test 
    - npm test
  
  - Cypress automation test:
    - cd automation-test
    - npm install
    - npm run dev


2. CURL: 
  - getTodo: curl -X GET http://localhost:5050 tasks -H "Authorization: testabc.xyz.ahk"
  

3. My solution is based on simple structure, it does not depend on the library. During development it will be easy to find related files.
Source code has consistent file and directory naming rules, based on a directory tree. The functions covered by automation test and unit test are quite suitable


4. If I had more time, I would handle the error when calling the api and show on screen in the most detail. I will use docker to create a database and query with real data and use jwt for login.
