FROM golang:alpine

WORKDIR /code
COPY build /code
CMD ["./build"]