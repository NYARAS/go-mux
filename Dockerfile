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
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server github.com/NYARAS/go-mux

FROM debian:bullseye-slim AS production

# Install necessary libraries for wkhtmltopdf.
RUN apt-get update && apt-get install -y --no-install-recommends \
    wkhtmltopdf \
    libstdc++6 \
    libx11-6 \
    libxrender1 \
    libxext6 \
    libfontconfig1 \
    fonts-dejavu \
    fonts-droid-fallback \
    fonts-freefont-ttf \
    fonts-liberation \
    libqt5webkit5 \
    libqt5widgets5 \
    libqt5gui5 \
    libqt5core5a \
    libqt5network5 \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /server

WORKDIR /app

CMD ["/server"]
