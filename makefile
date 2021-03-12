psql-up:
	docker-compose up -d
	echo "Wait 5s for postgres up and running..."
	sleep 5s
	PGPASSWORD=phuonghau createdb -U phuonghau -h localhost -p 8899 togo
	PGPASSWORD=phuonghau psql -U phuonghau -p 8899 -h localhost -f data_seed.sql togo
psql-down:
	docker-compose down