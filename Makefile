.PHONY: build up down
build:
	docker-compose build
up:
	docker-compose up backend && docker-compose rm -fsv
down:
	docker-compose down --volumes

test:
	docker-compose up test && docker-compose rm -fsv

binary:
	env GOOS=linux GOARCH=386 go build -o build .

test-binary:
	env GOOS=linux GOARCH=386 go test -c testing

clean:
	docker rm -f $(docker ps -a -q)