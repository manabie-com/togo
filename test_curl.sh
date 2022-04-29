#!/bin/bash
curl --cookie "user_id=$1" http://localhost:8090/tasks  -X POST -H 'Content-Type: application/json' -d "{\"title\": \"$2\", \"content\": \"$3\"}"