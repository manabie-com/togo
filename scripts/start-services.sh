set -a
source .env

sh ./scripts/mongo/mongo-start.sh
sh ./scripts/kafka/kafka-start.sh
sh ./scripts/redis/redis-start.sh
