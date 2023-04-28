# FROM golang:1.18 as builder

# WORKDIR /app

# COPY . .

# RUN go mod tidy

# RUN CGO_ENABLED=0 go build -o main ./cmd/http/main.go

# FROM alpine:3

# WORKDIR /app

# COPY --from=builder /app .

# EXPOSE 8080

# CMD ["./main"]

FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/http/main.go

EXPOSE 8080

CMD ["./main"]