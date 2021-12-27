default:
	@echo "###################-Marvik Challenge MakeFile-##########################"
	@echo "- build: Build the service"
	@echo "- devshell: Run the container and open a sh shell inside it"
	@echo "- up: Run the container and executes the service inside"
	@echo "- down: Take down the service container and all the dependent containers"
	@echo "###################-Marvik Challenge MakeFile-##########################"

init:
	@cp .env.template .env

build:
	@docker-compose build --no-cache

devshell:
	@docker-compose run --rm --service-ports api sh

up:
	@docker-compose run --rm --service-ports api

down:
	@docker-compose down --remove-orphans

start:
	@go run main.go

database: 
	@docker exec -it marvik_db_1 bash -c "psql -U gorm" 

