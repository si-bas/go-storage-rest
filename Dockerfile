# Builder
FROM golang:1.20-alpine3.17 as builder

RUN apk update && apk upgrade

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build main.go

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/main /app
COPY --from=builder /app/.env /app

ENTRYPOINT ["sh", "-c", "/app/main serve"]