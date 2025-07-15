##FROM golang:1.13-alpine3.11 AS build
##RUN apk --no-cache add gcc g++ make ca-certificates
##WORKDIR /go/src/github.com/akhilsharma90/go-graphql-microservice
##COPY go.mod go.sum ./
##COPY vendor vendor
##COPY account account
##RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./account/cmd/account
##
##FROM alpine:3.11
##WORKDIR /usr/bin
##COPY --from=build /go/bin .
##EXPOSE 8080
##CMD ["app"]
#
## Stage 1: Build
#FROM golang:1.20-alpine AS build
#WORKDIR /app
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app/myapp
#
## Stage 2: Run
#FROM alpine:latest
#WORKDIR /app
#
#COPY --from=build /app/myapp /app/myapp
#
#EXPOSE 8080
#CMD ["/app/myapp"]


# Stage 1: Build
FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
COPY sendmail sendmail
COPY pkg pkg
COPY /sendmail/config config

RUN APP_ENV=docker CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./sendmail/cmd/sendmail

# Stage 2: Run
FROM alpine:latest

WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY sendmail/config /etc/sendmail/
EXPOSE 7000

CMD ["app"]
