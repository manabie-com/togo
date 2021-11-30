-To run project
npm install & npm start

Auth Credential
Username: firstUser
Password: example

-A sample “curl” command to call your API
curl http://localhost:3001/todos -H "Authorization: testabc.xyz.ahk"

-How to run your unit tests locally?
npm test
then 
find and click auth.spec.js under Authentication folder (test Auth flow)
find and click createTodo.spec.js under CreateTodo folder (test limit 5 task per user per day flow)

-What do you love about your solution?

My simple solution base on read write file using fs