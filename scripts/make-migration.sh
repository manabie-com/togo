#!/usr/bin/env bash

set -e

timestamp() {
  date +"%Y%m%d_%H%M%S"
}

touch ./migrations/"$(timestamp)"_"$name".sql