# Start from the latest Golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Daiki Nakashita <daiki815@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Download and install pressly/goose
RUN go install github.com/pressly/goose/cmd/goose@latest

# This container will be executable, so ensure goose is in the PATH
ENV PATH="/go/bin:${PATH}"
