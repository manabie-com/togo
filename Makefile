storages_test:
	docker-compose --env-file ./configs/app.test.env up storages_test

app:
	docker-compose --env-file ./configs/app.dev.env up app