up-local:
	docker-compose -f ./docker/docker-compose.yml --env-file .env up -d

down-local:
	docker-compose -f ./docker/docker-compose.yml down --remove-orphans