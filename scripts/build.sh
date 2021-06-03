#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"
MAIN_PATH="$(git rev-parse --show-toplevel || echo ".")/cmd/server/main.go"
BIN_PATH="$(git rev-parse --show-toplevel || echo ".")/bin/"
build() {
  cmd=$1
  echo_info "Building $cmd"
  target=$(echo $cmd | sed 's/cmd/bin/')
  go build -ldflags="$GO_LDFLAGS" -i -v -o $target $cmd
  ls -lah -d $target
}

build_all() {
  # Build
  for cmd_package in $(find ./cmd -type d); do
    # skip the folder if there's no go file in it
    ls "$cmd_package"/*.go >/dev/null 2>&1 || continue
    # build the cmd
    build "$cmd_package"
  done
}

build_linux() {
  env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o $BIN_PATH"/linux-server" $MAIN_PATH
}
