# syntax=docker/dockerfile:1
FROM golang:1.18.1-alpine3.15
WORKDIR /app
COPY ./ /app
RUN apk add --no-cache git && CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:3.15
WORKDIR /app
COPY --from=0 /app/app ./app
CMD ["./app"]
