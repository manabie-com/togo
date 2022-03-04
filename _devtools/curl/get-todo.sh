ACCESS_TOKEN=$(curl --location --request POST 'http://localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "test01@simple.app",
  "password": "testing123"
}' | grep -Po '"'"token"'"\s*:\s*"\K([^"]*)')
echo '--'
curl --location --request GET "http://localhost:3000/todos/621de43cd5bb8c03e0643107" \
  --header "Authorization: Bearer $ACCESS_TOKEN" 
echo ""