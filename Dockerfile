# Dockerfile References:

# Start from the latest golang base image
FROM golang:1.14

# Maintainer Info
LABEL maintainer="Jeffrey Lane <jeffrey.lane@nielsen.com>"

# Set the current working directory inside the container
WORKDIR /go/src/signals

ENV GOPATH /go/src
ENV GOOS linux
ENV GOARCH amd64
# Copy the source from the current directory to the Working Dir inside the container
COPY . .

# Expose the port 8082 to the outside world
#EXPOSE 8082
RUN chmod ug+x main

# Define the entrypoint
ENTRYPOINT [ "/go/src/signals/main" ]