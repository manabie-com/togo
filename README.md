### Requirements
.net core 6 sdk/ .net core 6 runtime
### How to run source code locally: 

after installing the above requirement, use command
go to src

> $ dotnet build

go to bin of Manabie.Api, use command
> $ dotnet Manabie.Api/bin/Debug/net6.0/Manabie.Api.dll  
### CURL: 
#### Authenticate
> curl -X 'POST' \
> 'https://localhost:5001/api/Auth/authenticate' \
> -H 'accept: */*' \
> -H 'Content-Type: application/json' \
> -d '{
> "username": "cuongnsm",
> "password": "password"
> }'
####	Authorization test
>  curl -X 'GET' \
>  'https://localhost:5001/api/Auth' \
>  -H 'accept: */*' \
>  -H 'Authorization: Bearer {token}'
#### Add task
>   curl -X 'POST' \
>   'https://localhost:5001/api/Task' \
>   -H 'accept: text/plain' \
>   -H 'Content-Type: application/json' \
>   -H 'Authorization: Bearer {token}'
>   -d '{
>   "todo": "string"
> }'

#### Get tasks

> curl -X 'GET' \
>   'https://localhost:5001/api/Task' \
>   -H 'accept: text/plain' \
> -H 'Authorization: Bearer {token}'
#### Run tests
To run test, go to src following use command
> $ dotnet test
#### LOVE!
##### I do love simple!
