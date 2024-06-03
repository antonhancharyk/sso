FROM golang:alpine AS builder

WORKDIR /opt/app

COPY . .

RUN go build -o sso ./cmd/main.go

FROM alpine

WORKDIR /opt/app

COPY --from=builder /opt/app/sso /opt/app/sso

EXPOSE 8080

CMD ["./sso"]