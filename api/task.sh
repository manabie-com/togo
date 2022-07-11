#!/usr/bin/env bash -l

. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"
. "$(git rev-parse --show-toplevel || echo ".")/api/.env"

cmd=$1
token="Bearer "$token

add_task() {
  id=${1:-1}
  title=${2:-task_title}
  description=${3:-task_description}
  priority=${4:-1}
  http -vv POST  $API_URL/tasks user_id:=$id title=$title description=$description priority:=$priority
}

"$@"
