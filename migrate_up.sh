#!/bin/bash
folder=$(basename "$PWD")
echo folder
docker-compose down
docker-compose -f docker-compose.migrate.yaml up -d
docker exec -it $folder-migrations-1 bash -c 'npx migrate up'
docker-compose down