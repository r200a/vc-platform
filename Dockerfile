# Stage 1 - Build the binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# copy dependency files first (better docker cache)
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build the binary 
RUN go build -o server ./cmd/main.go

# Stage 2 - run the binary in a minimal image 
FROM alpine:latest

WORKDIR /app

# copy only the binary from builder stage
COPY --from=builder /app/server .

EXPOSE 8085

CMD ["./server"]

