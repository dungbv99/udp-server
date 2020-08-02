FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main server/server.go
CMD ["./main"]