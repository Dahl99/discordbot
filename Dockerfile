# Builder container
FROM golang:1.21 AS builder

# Set working directory
WORKDIR /src

# Update and install apt dependencies
RUN apt-get update && apt-get install -q -y stockfish

# Copy go module files
COPY go.mod go.sum ./

# Download go module dependencies
RUN go mod download all

# Copy source files and build executable
COPY . ./
RUN go build -v -o /bin/app ./cmd/discordbot/main.go

# Production container
FROM debian:bullseye-slim AS final

# Set working directory
WORKDIR /app

# Update and install apt dependencies
RUN apt-get update && \
    apt-get install -q -y \
    curl \
    python3 \
    ffmpeg \
    inkscape

# Symlink python3 to python
RUN ln -s /usr/bin/python3 /usr/bin/python

# Copy stockfish from builder and add to path
COPY --from=builder /usr/games/stockfish /usr/games
ENV PATH /usr/games:$PATH

# Copy config.json file
COPY .env ./

# Copy executable from builder
COPY --from=builder /bin/app ./

# Start app
CMD ["./app"]