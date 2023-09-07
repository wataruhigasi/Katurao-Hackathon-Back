u:
	docker-compose down --remove-orphans && docker-compose up -d

build:
	docker-compose build

d:
	docker-compose down --remove-orphans

log:
	docker-compose logs -f

ps:
	docker-compose ps