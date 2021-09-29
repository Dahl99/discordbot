FROM golang:1.17.1

# Set working directory
WORKDIR /app

# Copy go module files and config for bot
COPY go.mod .
COPY go.sum .
COPY config.json .

# Update and install apt dependencies
RUN apt update && apt install -y \
    python3 \
    ffmpeg \
    nodejs \
    npm

# Install nodemon
RUN npm install -g nodemon

# Symlink python3 to python
RUN ln -s /usr/bin/python3 /usr/bin/python

# Download go module dependencies
RUN go mod download all

# Download youtube-dl and chmod correct permissions
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

# Start app using nodemon
CMD ["nodemon", "--exec", "go", "run", "/app/main.go", "--signal", "SIGTERM"]