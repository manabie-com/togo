PROJECT=simple-app
docker rmi todo_svc:testing

docker-compose -p $PROJECT -f ./docker-compose-test.yml down --volumes
