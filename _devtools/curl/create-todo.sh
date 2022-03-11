ACCESS_TOKEN=$(curl --location --request POST 'http://localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "test01@simple.app",
  "password": "testing123"
}' | grep -Po '"'"token"'"\s*:\s*"\K([^"]*)')
echo '--'
# echo $ACCESS_TOKEN  
curl --location --request POST 'http://localhost:3000/todos' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $ACCESS_TOKEN" \
--data-raw '{
  "id": "621de43cd5bb8c03e0643107",
  "userId": "55f01d19-c913-4b0c-ba0e-63ead676b0d5",
  "title": "Do task 01",
  "desc": "test01@simple.app"
}' 
echo ""