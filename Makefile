.PHONY: build up down

create_migration:
	goose -dir ./migrations create $(NAME)

migrate:
	goose -dir ./migrations up

create-dev-db:
	docker exec -it graphyy-postgresql-dev1 psql -U postgres -c "CREATE DATABASE graphyy_development;"

drop-dev-db:
	docker exec -it graphyy-postgresql-dev1 psql -U postgres -c "DROP DATABASE graphyy_development;"

build:
	env GOOS=linux GOARCH=386 go build -o build ./cmd/server/main.go
	docker-compose build
run-db:
	docker-compose up -d postgresql-dev
up:
	env GOOS=linux GOARCH=386 go build -o build ./cmd/server/main.go
	docker-compose up backend && docker-compose rm -fsv
down:
	docker-compose down --volumes

test:
	docker-compose -f docker-compose.test.yml build test
	docker-compose -f docker-compose.test.yml up -d test
	docker-compose -f docker-compose.test.yml exec -T test ginkgo -r $(filter-out $@,$(MAKECMDGOALS))
	docker-compose -f docker-compose.test.yml rm -fsv

clear-test:
	docker volume remove graphyy_postgres_test_data

binary:
	env GOOS=linux GOARCH=386 go build -o build ./cmd/server/main.go

clean-containers:
	docker rm -f $(docker ps -a -q)

clean-images:
	docker image prune
