FROM golang:alpine

WORKDIR /code
COPY . .
RUN go mod download

RUN env GOOS=linux GOARCH=386 go build -o build .

CMD ["./build"]