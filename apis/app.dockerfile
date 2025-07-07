# Stage 1: Build
FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
COPY apis apis
COPY pkg pkg
COPY /apis/config config

RUN APP_ENV=docker CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./apis/cmd/apis

# Stage 2: Run
FROM alpine:latest

WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY apis/config /etc/apis/
EXPOSE 6002

CMD ["app"]
