FROM golang:1.17-bullseye AS builder

# Create and change to the app directory.
WORKDIR /app

# Copy go.sum/go.mod and warm up the module cache.
COPY go.* ./
RUN go mod download

# Set the environment variable for Gin in release mode.
ENV GIN_MODE release

# Now copy the rest of the application's source code
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/server github.com/NYARAS/go-mux

FROM debian:bullseye-slim AS production

COPY --from=builder /app/server /server

ENTRYPOINT ["./server"]
