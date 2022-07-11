#!/usr/bin/env bash -l

. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"
. "$(git rev-parse --show-toplevel || echo ".")/api/.env"

cmd=$1
token="Bearer "$token

create_one() {
  name=${1:-dat}
  email=${2:-datshiro@gmail.com}
  quota=${3-5}
  http -vv $API_URL/users name=$name email=$email quota:=$quota
}

get_one() {
  id=${1:-1}
  http -vv  $API_URL/users/$id
}


update_one() {
  id=${1:-1}
  name=${2:-dat}
  email=${3:-datshiro@gmail.com}
  password=${4:-123456789}
  http -vv PUT  $API_URL/users/$id Authorization:"$token" name=$name email=$email password=$password
}

add_task() {
  id=${1:-1}
  title=${2:-task_title}
  description=${3:-task_description}
  priority=${4:-priority}
  http -vv POST  $API_URL/tasks/$id  title=$title description=$description priority=$priority
}

"$@"
