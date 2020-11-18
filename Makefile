.PHONY: build up down
build:
	docker-compose build
up:
	docker-compose up && docker-compose rm -fsv
down:
	docker-compose down --volumes
test:
	docker-compose run backend env GOOS=linux GOARCH=386 go test -v .
	# docker exec -it backend env GOOS=linux GOARCH=386 go test -v .

go-build:
	env GOOS=linux GOARCH=386 go build -o build .

clean:
	docker rm -f $(docker ps -a -q)