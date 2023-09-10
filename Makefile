u:
	docker-compose down --remove-orphans && docker-compose up -d

b:
	docker-compose up --build

d:
	docker-compose down --remove-orphans

r:
	docker volume rm katurao-hackathon-back_db-data

init:
	docker-compose down --remove-orphans && docker volume rm katurao-hackathon-back_db-data && docker-compose up --build -d