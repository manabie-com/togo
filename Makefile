storages_test:
	docker-compose --env-file ./configs/app.test.env up storages_test

usecase_test:
	docker-compose --env-file ./configs/app.test.env up usecase_test

integration_test:
	docker-compose --env-file ./configs/app.test.env up integration_test

dev:
	docker-compose --env-file ./configs/app.dev.env up dev