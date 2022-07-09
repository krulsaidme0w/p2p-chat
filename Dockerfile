FROM golang:1.17-buster AS build

WORKDIR /app
ADD . .

ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build ./cmd/main.go

FROM alpine:latest

RUN apk upgrade --update-cache --available && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build /app/ .
ADD cmd .

EXPOSE 25042

CMD ["./main"]