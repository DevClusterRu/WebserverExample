FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main
FROM alpine:3
COPY --from=builder /app/main /bin/main
EXPOSE 9999
ENTRYPOINT ["/bin/main"]

