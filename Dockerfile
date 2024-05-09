FROM golang:1.18 as builder
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
FROM alpine:latest
RUN adduser -D myuser
USER myuser
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]