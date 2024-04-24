FROM golang:alpine

WORKDIR /code
COPY build /code

# Run the binary program produced by `go build`
CMD ["./build"]