curl --location --request POST 'http://localhost:3000/limit-settings' \
--header 'Content-Type: application/json' \
--header "Authorization: Bearer $ACCESS_TOKEN" \
--data-raw '{
  "name": "TODO_LIMIT",
  "value": 3,
  "userId": "55f01d19-c913-4b0c-ba0e-63ead676b0d5"
}'
echo ""