FROM golang:1.22.5-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

RUN mkdir -p /root/configs
COPY configs/local.yaml /root/configs/
COPY .env .

RUN mkdir -p /root/db/migrations
COPY db/migrations /root/db/migrations/

EXPOSE 44044

CMD ["./main"]