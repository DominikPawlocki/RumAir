#FROM alpine:3.5 golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore
# -----------------------------------
FROM golang:alpine AS builder

# Add all the source code (except what's ignored# under `.dockerignore`) to the build context.
COPY ./ /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API/

WORKDIR /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API/
RUN go build -o /usr/bin/RumAir_Pmpro_Sensors_API

# RUN set -ex && \
#   cd /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API && \       
#   CGO_ENABLED=0 go build \
#         -tags netgo \
#         -v -a \
#         -ldflags '-extldflags "-static"' && \
#   mv ./RumAir_Pmpro_Sensors_API /usr/bin/RumAir_Pmpro_Sensors_API

#swagger.json generation from a code, for swaggerUI serving 
FROM quay.io/goswagger/swagger as swaggerspecbuilder
COPY --from=builder /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API /usr/local/bin/RumAir_Pmpro_Sensors_API

WORKDIR /usr/local/bin/RumAir_Pmpro_Sensors_API
# RUN set -ex && \
#   cd /usr/local/bin/RumAir_Pmpro_Sensors_API && \
#     swagger generate spec -o ./swagger.json
RUN set -ex && \
  swagger generate spec -o ./swagger.json

# validate a generated file
RUN set -ex && \
    swagger validate --stop-on-error ./swagger.json 
RUN pwd
#copy generated artifact (swagger.json) from container filesystem into host filesystem to be used later by docker-compose swaggerUI component
#COPY ./usr/local/bin/RumAir_Pmpro_Sensors_API/swagger.json /usr/local/bin 

#last FROM statement is the final base image.
FROM busybox

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/RumAir_Pmpro_Sensors_API /usr/local/bin/RumAir_Pmpro_Sensors_API

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "RumAir_Pmpro_Sensors_API" ]




# LABEL "com.example.vendor"="ACME Incorporated"
# LABEL com.example.label-with-value="foo"
# LABEL version="1.0"
# LABEL description="This text illustrates \
# that label-values can span multiple lines."