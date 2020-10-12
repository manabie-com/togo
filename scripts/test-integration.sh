#!/bin/bash
source ./internal/config/.env.sh && \
cd ./internal/integration && \
go test -timeout 300s -cover -a -v