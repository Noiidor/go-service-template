FROM golang:1.23-alpine3.20 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -v -o /plain-http ./cmd/plain-http

FROM alpine:3.20 as run

COPY --from=build /plain-http /plain-http

WORKDIR /app
CMD ["/plain-http"]
