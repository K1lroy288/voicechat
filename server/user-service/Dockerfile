FROM golang:alpine AS builder

WORKDIR /user-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /user-service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /user-service/user-service .

COPY .env .

EXPOSE 8000

CMD ["./user-service"]
