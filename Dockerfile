FROM golang:1.23-alpine3.20 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -v -o /wizards-rest ./cmd/wizards-rest

FROM alpine:3.20 as run

COPY --from=build /wizards-rest /wizards-rest

WORKDIR /app
CMD ["/wizards-rest"]
