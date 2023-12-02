FROM golang:1.20.3 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN adduser -D nonrootuser

WORKDIR /root/
COPY --from=builder /app/server .
COPY ./cmd/server/config.local.yaml /app/config/
COPY ./cmd/server/config.canary.yaml /app/config/
COPY ./cmd/server/config.production.yaml /app/config/

ENV USER_ACCOUNT_CONFIG_PATH=/app/config

EXPOSE 8080

USER nonrootuser
CMD ["./server"]
