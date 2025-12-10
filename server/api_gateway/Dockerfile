FROM golang:alpine AS builder

WORKDIR /api-gateway

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-gateway

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /api-gateway/api-gateway .

COPY .env .

EXPOSE 8000

CMD ["./api-gateway"]
