#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

./bin/golangci-lint run --fast $*

EXIT_CODE=$?
exit ${EXIT_CODE}
