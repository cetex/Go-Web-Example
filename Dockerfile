# Get the latest golang docker-image
FROM golang:latest
# Set WORKDIR (Initial CWD for all commands in image) to /app
WORKDIR /app

# Copy everything to WORKDIR
COPY . ./
# Download go module dependencies (if any)
RUN go mod download
# Build go container
RUN go build ./

# The webserver runs on port 8080, propose a portforward from the host to the container
EXPOSE 8080/tcp

# Set default command to "/app/main" so we don't need to know the binary name when running the docker-image, it will default to this command.
CMD /app/main
