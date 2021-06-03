#!/usr/bin/env bash

# This is shortcut for running `migrate` against the default database as
# configured in config/database.json

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"
. "$(git rev-parse --show-toplevel || echo ".")/.conf"

cd "$PROJECT_DIR"
cmd=$1
shift
echo_info "Run migrate command: $cmd"
./bin/migrate -verbose -database "$DB_URL" -path ./db/migrations/ $cmd $*
cd $WORKING_DIR
