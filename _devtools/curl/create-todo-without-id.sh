echo 'It will try to create 4 todo.'
sleep 5

ACCESS_TOKEN=$(curl --location --request POST 'http://localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "test01@simple.app",
  "password": "testing123"
}' | grep -Po '"'"token"'"\s*:\s*"\K([^"]*)')
echo '--'


curl --location --request POST 'http://localhost:3000/todos' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $ACCESS_TOKEN" \
--data-raw '{
  "userId": "55f01d19-c913-4b0c-ba0e-63ead676b0d5",
  "title": "Do task 01",
  "desc": "test01@simple.app"
}'
echo ''

echo 'If we already invoked create-todo the two below requests will respond with error. Otherwise only last request will fail.'
curl --location --request POST 'http://localhost:3000/todos' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $ACCESS_TOKEN" \
--data-raw '{
  "userId": "55f01d19-c913-4b0c-ba0e-63ead676b0d5",
  "title": "Do task 01",
  "desc": "test01@simple.app"
}'
echo ''

curl --location --request POST 'http://localhost:3000/todos' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $ACCESS_TOKEN" \
--data-raw '{
  "userId": "55f01d19-c913-4b0c-ba0e-63ead676b0d5",
  "title": "Do task 01",
  "desc": "test01@simple.app"
}'
echo ''

echo "Below request will be responded with error messsage"
curl --location --request POST 'http://localhost:3000/todos' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $ACCESS_TOKEN" \
--data-raw '{
  "userId": "55f01d19-c913-4b0c-ba0e-63ead676b0d5",
  "title": "Do task 01",
  "desc": "test01@simple.app"
}'
echo ""