FROM golang:1.25-alpine AS builder
WORKDIR /srv
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go build -tags dist -o ./sem1app ./cmd/main.go

FROM alpine:3.22 AS prod
COPY --from=builder /srv/sem1app /srv/docker-entrypoint.sh /usr/local/bin/
CMD ["docker-entrypoint.sh"]
