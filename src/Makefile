SHELL := /bin/bash
.EXPORT_ALL_VARIABLES:

HOST_ENV=127.0.0.1
USER_ENV=user
PASS_ENV=example

build:
	go build -o ./bin

dev-up:
	docker-compose -f ../docker/docker-compose.dev.yml up -d

dev-down:
	docker-compose -f ../docker/docker-compose.dev.yml down	

dev: dev-down dev-up build
	./bin/tool-checker

qa: qa-down
	docker-compose -f ../docker/docker-compose.qa.yml build
	docker-compose -f ../docker/docker-compose.qa.yml up -d

qa-down:
	docker-compose -f ../docker/docker-compose.qa.yml down

remove:
	docker rm -f $$(docker ps -aq)
