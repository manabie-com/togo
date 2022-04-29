#!/bin/bash
if [[ -z $1 ]]
then
    echo 'Missing migration-name. Usage ./migrate_create.sh [migration-name]'
    exit 1
fi


folder=$(basename "$PWD")
docker-compose down
docker-compose -f docker-compose.migrate.yaml up -d
docker exec -it $folder-migrations-1 bash -c "npx migrate create $1"
docker-compose down