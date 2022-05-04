set -a
source .env

sh ./scripts/mongo/mongo-down.sh
sh ./scripts/kafka/kafka-down.sh
sh ./scripts/redis/redis-down.sh
