
# Use an official Go runtime as a parent image
FROM golang:1.19.5

# Set the working directory in the container
WORKDIR /golang-assignment

# Copy the local source code to the container
COPY . .

# Build the Go application
RUN go build -o golang-assignment

# Expose a port the application will listen on
EXPOSE 8080

# Copy the configuration file to the container
COPY config.yaml /golang-assignment/config.docker.yaml 

ENTRYPOINT ./golang-assignment -c config.docker.yaml