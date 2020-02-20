#FROM alpine:3.5 golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore
# -----------------------------------
FROM golang:alpine AS builder

# Add all the source code (except what's ignored# under `.dockerignore`) to the build context.
ADD ./ /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API/

RUN set -ex && \
  cd /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./RumAir_Pmpro_Sensors_API /usr/bin/RumAir_Pmpro_Sensors_API

#last FROM statement is the final base image.
FROM busybox

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/RumAir_Pmpro_Sensors_API /usr/local/bin/RumAir_Pmpro_Sensors_API

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "RumAir_Pmpro_Sensors_API" ]