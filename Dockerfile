FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p /app/logs

RUN /go/bin/swag init --generalInfo cmd/main.go --output docs

RUN go build -o /app/main ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

RUN chmod +x ./main

COPY --from=builder /app/config /config


# Copy logs directory if it exists
#COPY --from=builder /app/logs /logs

EXPOSE 8080

CMD ["./main"]