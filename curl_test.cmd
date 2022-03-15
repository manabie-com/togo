
echo "=============>"
echo "=============>"
echo "\n-------------------- HealthCheck ------------------- \n"
curl --location --request GET 'http://localhost:9099/api/v1/healthcheck/status'
echo "\n-------------------- END -------------------"


echo "\n-------------------- Register ------------------- \n"
curl --location --request POST 'http://localhost:9099/api/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"dio123",
    "password":"dio123"
}'
echo "\n-------------------- END -------------------"


echo "\n-------------------- LOGIN ------------------- \n"
curl --location --request POST 'http://localhost:9099/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"dio123",
    "password":"dio123"
}'
echo "\n-------------------- END -------------------"


echo "\n-------------------- Create Todo ------------------- \n"
curl --location --request POST 'http://localhost:9099/api/v1/todo' \
--header 'Authorization: 6b2ffd41-4419-41a7-8504-bf3d68293f48' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content":"a 22"
}'
echo "\n-------------------- END -------------------"


echo "\n-------------------- Update Todo ------------------- \n"
curl --location --request PUT 'http://localhost:9099/api/v1/todo' \
--header 'Authorization: 8dded146-3c79-408c-9760-32189ab45f51' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "content": "xyz",
    "status": 1
}'
echo "\n-------------------- END -------------------"


echo "\n-------------------- Get Todo ------------------- \n"
curl --location --request GET 'http://localhost:9099/api/v1/todo?size&index' \
--header 'Authorization: b2298c9b-65c7-49b7-92b2-44a9bf78b436'
echo "\n-------------------- END -------------------"