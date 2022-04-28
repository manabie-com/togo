<h1 align="center" id="title">ToGo Project</h1>
<h3 align="center" id="title">By: Kier A. Quebral</h3>

<h2>Kindly use the follwing command:</h2>

<p>1. To run the pprogram in your localhost</p>

```
make run
```

<p>2. To kill the current process</p>

```
make kill
```

<p>3. To do a unit testing</p>

```
make testing.unit
```

<h2>🧐 Here're some of the project's best features:Use this cURL to call our awesome API</h2>

<p>1. Use this cURL to create you awesome user.</p>

```
curl --location --request POST 'http://127.0.0.1:8080/api/v1/user/create' \ --header 'Content-Type: application/json' \ --data-raw '{     "user_name":"admin_kier"     "password":"P@ssw0rd123"     "max_todos":5 }'
```

<p>2. Use this cURL to login your awesome account.. Once you successfully login it will return a token. You need the token to create task</p>

```
curl --location --request GET 'http://127.0.0.1:8080/api/v1/user/login' \ --header 'Content-Type: application/json' \ --data-raw '{     "user_name":"admin_kier"     "password":"P@ssw0rd123" }'
```

<p>3. Use this cURL to create your awesome task. Remember you need to include the token in Authorization header</p>

```
curl --location --request POST 'http://127.0.0.1:8080/api/v1/task/create' \ --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjUxMzAyNTE4LCJpZGVudGl0eSI6ImFkbWluX2tpZXIifQ._0VN84kMSuLFgBzZRwRoLVlVrmj_2GCKEJQ4t2bjlnE' \ --header 'Content-Type: application/json' \ --data-raw '{     "name":"Tas"     "description": "Lorem Ipsum amet"     "user_id":1 }'
```

<p>What I love for my solution is I show how to use interface and I share how I write a simple api that use jwt for authentication. I hope this project met your requirements for the position as Back-End Engineer. </p>

