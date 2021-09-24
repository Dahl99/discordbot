FROM --platform=linux/amd64 golang:1.16

# Set working directory
WORKDIR /app

# Copy go moduel files and config for bot
COPY go.mod .
COPY go.sum .
COPY config.json .

# Update and install apt dependencies
RUN apt update
RUN apt install -y python3
RUN apt install -y ffmpeg

# Symlink python3 to python
RUN ln -s /usr/bin/python3 /usr/bin/python

# Download go module dependencies
RUN go mod download

# Download youtube-dl and chmod correct permissions
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

# Install cosmtrek/air for live reloading
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Start app with cosmtrek/air
CMD ["air"]