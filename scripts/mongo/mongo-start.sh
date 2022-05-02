echo "MongoDB Service - Start..."

MONGO_DATA_PATH="./scripts/mongo/data/"

if [ ! -d "$MONGO_DATA_PATH" ] || [ -z "$(ls -A $MONGO_DATA_PATH)" ];
then
  mkdir $MONGO_DATA_PATH
  chmod +x $MONGO_DATA_PATH
fi

docker-compose -f ./scripts/mongo/docker-compose.yaml up -d