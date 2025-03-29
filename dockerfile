FROM golang:1.24.1-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /bin/app ./cmd/server

FROM alpine:latest

RUN apk --no-cache add sqlite

WORKDIR /root/

COPY --from=builder /bin/app /bin/app

EXPOSE 8080

CMD ["/bin/app"]
