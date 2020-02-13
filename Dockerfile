#FROM alpine:3.5 golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore

# -------------------------------------------
# WORKDIR /go/src/app

# COPY . .

# RUN go get -d -v ./...
# # or go install might be faster in that case.
# RUN go build -v ./...

# -----------------------------------
FROM ubuntu:latest AS builder

# Add all the source code (except what's ignored# under `.dockerignore`) to the build context.
ADD ./ /go/src/github.com/DominikPawlocki/RumAir/

ENV RUMAIR_DATABASE = 'hello'
ENV RUMAIR_DATABASE_PASSWORD = 'aaaa'

# RUN sudo apt-get install ca-certificates

RUN set -ex && \
  cd /go/src/github.com/DominikPawlocki/RumAir && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./RumAir /usr/bin/RumAir

# FROM busybox BusyBox combines tiny versions of many common UNIX utilities into a single small executable.

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/RumAir /usr/local/bin/RumAir

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "RumAir" ]