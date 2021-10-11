run-docker-compose: 
	docker-compose up -d
kill-port: 
	sudo kill -9 `sudo lsof -t -i:5510`