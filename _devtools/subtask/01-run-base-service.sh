
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

PROJECT=simple-app
# Deploy core services
docker-compose -p $PROJECT -f ./docker-compose-test.yml up -d
sleep 5