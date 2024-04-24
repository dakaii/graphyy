.PHONY: build up down

create_migration:
	goose -dir ./migrations create $(NAME)

migrate:
	goose -dir ./migrations up

build:
	docker-compose build
up:
	docker-compose up backend && docker-compose rm -fsv
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
