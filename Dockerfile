FROM golang:1.17.1 AS builder

# Set working directory
WORKDIR /src

# Copy go module files and config for bot
COPY go.mod go.sum ./

# Download go module dependencies
RUN go mod download all

COPY . ./
RUN go build -v -o /bin/app ./*.go

FROM debian:buster-slim

WORKDIR /app

# Update and install apt dependencies
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y \
    curl \
    python3 \
    ffmpeg

# Symlink python3 to python
RUN ln -s /usr/bin/python3 /usr/bin/python

# Download youtube-dl and chmod correct permissions
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

COPY config.json ./
COPY --from=builder /bin/app ./

# Start app using nodemon
CMD ["./main"]