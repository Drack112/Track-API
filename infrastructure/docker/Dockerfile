FROM golang:1.25.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOOS=linux GOARCH=amd64 go build -o track-api ./cmd/track-api

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /app/track-api .
COPY infrastructure/docker/scripts/entrypoint.sh .
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoiont.sh"]
