#FROM alpine:3.5 golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore
# -----------------------------------
FROM golang:alpine AS builder

# Add all the source code (except what's ignored# under `.dockerignore`) to the build context.
ADD ./ /go/src/github.com/DominikPawlocki/RumAir/

ENV RUMAIR_DATABASE = 'hello'
ENV RUMAIR_DATABASE_PASSWORD = 'aaaa'

RUN set -ex && \
  cd /go/src/github.com/DominikPawlocki/RumAir && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./RumAir /usr/bin/RumAir

#last FROM statement is the final base image.
FROM busybox

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/RumAir /usr/local/bin/RumAir

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "RumAir" ]