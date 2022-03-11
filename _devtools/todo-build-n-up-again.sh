# export DOCKER_HOST=ssh://ubuntu@192.168.1.98

export MONGO_INITDB_ROOT_USERNAME=admin
export MONGO_INITDB_ROOT_PASSWORD=admin123
export MONGO_HOST=mongodb
export MONGO_PORT=27017
export MONGO_USER=admin
export MONGO_PASS=admin123

export RABBITMQ_USER=admin
export RABBITMQ_PASS=admin
export RABBITMQ_HOST=rabbitmq
export RABBITMQ_PROTOCOL=amqp
export RABBITMQ_PORT=5672

docker stop togo_todo_svc_1
docker rm stop togo_todo_svc_1


docker-compose -f ./docker-compose-dev.yml up -d --build todo_svc
