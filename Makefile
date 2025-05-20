.PHONY: run test swagger docker-build up-build \
        up-core up-all down logs rebuild

## ----------------------------
##	APP Commands
## ----------------------------
run:
	go run cmd/main.go

test:
	go test ./...

swagger:
	swag init --parseDependency --parseInternal -g cmd/main.go




## ----------------------------
##	Docker APP Commands
## ----------------------------

docker-build:
	docker build . -t task-tracker:latest

up-build:
	docker-compose -p task-tracker-app up -d




## ----------------------------
##  Infrastructure (infra-configs) Commands
## ----------------------------

COMPOSE_CORE=docker-compose -f infra-configs/docker-compose.core.yml
COMPOSE_LOGGING=$(COMPOSE_CORE) -f infra-configs/docker-compose.logging.yml

COMPOSE_CORE_NAME=-p core-task-tracker-app
COMPOSE_LOGGING_NAME =-p all-task-tracker-app

up-core: 
	$(COMPOSE_CORE) $(COMPOSE_CORE_NAME) up -d

up-all:
	$(COMPOSE_LOGGING) $(COMPOSE_LOGGING_NAME) up -d

down:
	$(COMPOSE_LOGGING) $(COMPOSE_LOGGING_NAME) down

logs:
	docker logs -f filebeat

reload:
	$(COMPOSE_LOGGING) $(COMPOSE_LOGGING_NAME) down -v
	$(COMPOSE_LOGGING) $(COMPOSE_LOGGING_NAME) up -d
