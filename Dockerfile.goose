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

# Copy the migration files into the container
COPY ./migrations /app/migrations

# Run the migrations when the container starts
# Replace the connection string with your actual database connection string
CMD goose -dir /app/migrations postgres "user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB host=$POSTGRES_HOST sslmode=disable" up