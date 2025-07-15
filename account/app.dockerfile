# Stage 1: Build
FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
COPY account account
COPY pkg pkg
COPY /account/config config

RUN APP_ENV=docker CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./account/cmd/account

# Stage 2: Run
FROM alpine:latest

WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY account/config /etc/account/ 
EXPOSE 8080

CMD ["app"]
