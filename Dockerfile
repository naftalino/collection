FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]