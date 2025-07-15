# Stage 1: Build
FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
COPY discord-logs-bot discord-logs-bot
COPY pkg pkg
COPY /discord-logs-bot/config config

RUN APP_ENV=docker CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./discord-logs-bot/cmd/discord-logs-bot

# Stage 2: Run
FROM alpine:latest

WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY discord-logs-bot/config /etc/discord-logs-bot/
EXPOSE 8080

CMD ["app"]
