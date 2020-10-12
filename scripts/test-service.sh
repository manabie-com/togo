#!/bin/bash
source ./internal/config/.env.sh && \
cd ./internal/services && \
go test -timeout 300s -cover -a -v