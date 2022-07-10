FROM golang:1.18-buster AS build

WORKDIR /app
ADD . .

ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build ./cmd/main.go

FROM sixeyed/ubuntu-with-utils

WORKDIR /app

COPY --from=build /app/ .
ADD cmd .

#ENTRYPOINT ["./main"]
CMD ["ping", "google.com"]