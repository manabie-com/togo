curl --location --request POST 'http://localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "test01@simple.app",
  "password": "testing123"
}'
echo ""