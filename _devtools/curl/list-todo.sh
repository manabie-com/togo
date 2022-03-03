ACCESS_TOKEN=$(curl --location --request POST 'http://localhost:3000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "test01@simple.app",
  "password": "testing123"
}' | grep -Po '"'"token"'"\s*:\s*"\K([^"]*)')
echo '--'

rawurlencode() {
  local string="${1}"
  local strlen=${#string}
  local encoded=""
  local pos c o

  for (( pos=0 ; pos<strlen ; pos++ )); do
     c=${string:$pos:1}
     case "$c" in
        [-_.~a-zA-Z0-9] ) o="${c}" ;;
        * )               printf -v o '%%%02x' "'$c"
     esac
     encoded+="${o}"
  done
  echo "${encoded}"    # You can either set a return variable (FASTER) 
  REPLY="${encoded}"   #+or echo the result (EASIER)... or both... :p
}
filter='
{
  "offset": 0,
  "limit": 100,
  "skip": 0,
  "include": [
    {
      "relation": "user",
      "scope": {
        "offset": 0,
        "limit": 100,
        "skip": 0
      }
    }
  ]
}
'
encode_qs=$( rawurlencode "$filter" )

curl --location --request GET "http://localhost:3000/todos?filter=$encode_qs" \
  --header "Authorization: Bearer $ACCESS_TOKEN" 
echo ""