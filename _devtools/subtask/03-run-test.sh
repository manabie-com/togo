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


NETWORK=simple_app_net_test
docker run \
  --rm \
  -it \
  --name todo_svc_testing \
  --network=$NETWORK \
  -e MONGO_HOST=$MONGO_HOST \
  -e MONGO_PORT=$MONGO_PORT \
  -e MONGO_USER=$MONGO_USER \
  -e MONGO_PASS=$MONGO_PASS \
  -e RABBITMQ_HOST=$RABBITMQ_HOST \
  -e RABBITMQ_PROTOCOL=$RABBITMQ_PROTOCOL \
  -e RABBITMQ_USER=$RABBITMQ_USER \
  -e RABBITMQ_PASS=$RABBITMQ_PASS \
  todo_svc:testing npm run test
