curl --location --request POST 'http://localhost:3000/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
  "id": "55f01d19-c913-4b0c-ba0e-63ead676b0d5",
  "username": "test01",
  "email": "test01@simple.app",
  "password": "testing123"
}'
echo ""