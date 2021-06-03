#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

cd "$PROJECT_DIR" || exit 1

# Prerequisite
#-------------------------------------------------------------------------------

# only complain about missing node, go or docker if we're not running on CI
if [[ -z $CI ]]; then
  if ! has go; then
    echo_error "Golang binary not found. Please install golang before continue"
    exit 1
  fi

  if ! has docker; then
    echo_warn "Docker not found. Please install or start it before running db"
  fi
fi

# Mandatory tools
#-------------------------------------------------------------------------------

echo_info "Download golang dependencies"
go get ./...

# create bin folder to store downloaded tools and compiled binaries
mkdir -p bin/

# Nice to have tools, should only be installed when not on CI, to save build time
#-------------------------------------------------------------------------------
if [[ -z $CI ]]; then
  if ! has ./bin/migrate; then
    echo_info "Install golang-migrate for database versioning"
    version="v4.12.2"
    if ! has apt-get; then
      platform="darwin"
      curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
      mv migrate.darwin-amd64 ./bin/migrate
    elif has apt; then
      platform="linux"
      curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
      mv migrate.linux-amd64 ./bin/migrate
    fi
    chmod +x ./bin/migrate
  fi

  if ! has ./bin/golangci-lint; then
    echo_info "Install golangci-lint for static code analysis (via curl)"
    # install into ./bin/
    # because different project might need different golang version,
    # and thus, need to use different linter version
    curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.23.7
  fi

  if ! has richgo; then
    echo_info "Install richgo for nicer go test output"
    # this one should be safe to install globally
    go get -v -u github.com/kyoh86/richgo
  fi

  if ! has goimports; then
    # this one is safe to install globally
    echo_info "Install goimports"
    go get -v -u golang.org/x/tools/cmd/goimports
  fi

  if ! has sqlboiler; then
    echo_info "Installing sqlboiler version 4. !"
    GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
    echo_info "Installing psql driver for sqlboiler "
    GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
  fi
fi
cp $PROJECT_DIR/.example.conf $PROJECT_DIR/.conf
cp $PROJECT_DIR/.example.env $PROJECT_DIR/.env

