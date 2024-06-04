FROM golang:alpine AS builder

WORKDIR /opt/app

COPY . .

RUN go build -o sso ./cmd/main.go

FROM alpine

WORKDIR /opt/app

COPY --from=builder /opt/app/sso /opt/app/sso

COPY --from=builder /opt/app/static /opt/app/static

COPY --from=builder /opt/app/migrations /opt/app/migrations

EXPOSE 8080

CMD ["./sso"]