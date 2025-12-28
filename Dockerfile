# Build stage
FROM golang:alpine AS build-stage
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app ./cmd/api

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Run smaller image
FROM alpine:3.23.2 AS build-release-stage
WORKDIR /root/
COPY --from=build-stage /usr/local/bin/app /usr/local/bin/app
EXPOSE 8080
CMD ["app"]
