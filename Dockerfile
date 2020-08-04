#FROM alpine:3.5 golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore
# -----------------------------------
FROM golang:alpine AS golang_builder

ARG BUILD_TYPE="manually triggered from docker build command, without compose"
RUN set -ex && echo "--------------- Build triggered by: $BUILD_TYPE -----------------"

# RUN $time=$(date +”%d-%b-%Y_%H:%M:%S”) && echo $time 

# Add all the source code (except what's ignored# under `.dockerignore`) to the build context.
COPY ./ /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API/

WORKDIR /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API/

# dont know why the simplified version doesnt work . Investigate !
#RUN go build -o /usr/bin/RumAir_Pmpro_Sensors_API

RUN set -ex && \
   cd /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API && \       
   CGO_ENABLED=0 go build \
         -tags netgo \
         -v -a \
         -ldflags '-extldflags "-static"' && \
   mv ./RumAir_Pmpro_Sensors_API /usr/bin/RumAir_Pmpro_Sensors_API

#-----------------------------------------------------------------------------------------------------
#swagger.json generation from a code, for swaggerUI serving 
FROM quay.io/goswagger/swagger as swagger_spec_builder
COPY --from=golang_builder /go/src/github.com/DominikPawlocki/RumAir_Pmpro_Sensors_API /usr/local/golangSourcesForSwagger

WORKDIR /usr/local/golangSourcesForSwagger
RUN set -ex && \
  swagger generate spec -o swagger.json

# validate a generated file and show whats inside directory
RUN set -ex && \
    swagger validate --stop-on-error swagger.json && \
    pwd && ls -lt
#-----------------------------------------------------------------------------------------------------

#last FROM statement is the final base image.
FROM busybox

# first copy the swagger.json from previous stage to the image for different docker container usage (swaggerui)
COPY --from=swagger_spec_builder /usr/local/golangSourcesForSwagger/swagger.json /tmp/swagger/

# Retrieve the binary from the previous stage
COPY --from=golang_builder /usr/bin/RumAir_Pmpro_Sensors_API /usr/local/bin/RumAir_Pmpro_Sensors_API

RUN cd /tmp/swagger && ls -lt

# Busybox is minimal. Really. Doesnt have 'apk' command even.. Need to copy it from previous stage
# COPY --from=golang_builder /etc/ssl/certs /etc/ssl/certs

ENV LISTEN_PORT=5000
EXPOSE 5000

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "RumAir_Pmpro_Sensors_API" ]



# LABEL "com.example.vendor"="ACME Incorporated"
# LABEL com.example.label-with-value="foo"
# LABEL version="1.0"
# LABEL description="This text illustrates \
# that label-values can span multiple lines."