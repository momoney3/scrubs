FROM golang:1.25

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

# # First stage: Build the application
# FROM golang:alpine AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o mycli
# # Second stage: Create a smaller image for the final executable
# FROM alpine
# WORKDIR /root/
# COPY --from=builder /app/mycli .
# ENTRYPOINT ["./mycli"]

