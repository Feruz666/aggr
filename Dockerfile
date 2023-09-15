# Build stage
FROM golang:1.20 AS builder

COPY . /go/src/app
WORKDIR /go/src/app
RUN go build -o client client/main.go && go build -o server server/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/app/client /app/
COPY --from=builder /go/src/app/server /app/
EXPOSE 8080
CMD ["./client", "&", "./server"]