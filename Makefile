.PHONY: build up down

create_migration:
	goose -dir ./migrations create $(NAME)

migrate:
	goose -dir ./migrations up

create-dev-db:
	docker exec -it postgresql-local psql -U postgres -c "CREATE DATABASE graphyy_development;"

drop-dev-db:
	docker exec -it postgresql-local psql -U postgres -c "DROP DATABASE graphyy_development;"


build:
	env GOOS=linux GOARCH=386 go build -o build ./cmd/server/main.go
	docker-compose build
run-db:
	docker-compose up -d postgresql
up:
	env GOOS=linux GOARCH=386 go build -o build ./cmd/server/main.go
	docker-compose up && docker-compose rm -fsv
down:
	docker-compose down --volumes

test:
	docker-compose up test && docker-compose rm -fsv

binary:
	env GOOS=linux GOARCH=386 go build -o build ./cmd/server/main.go

test-binary:
	env GOOS=linux GOARCH=386 go test -c testing

clean-containers:
	docker rm -f $(docker ps -a -q)

clean-images:
	docker image prune
