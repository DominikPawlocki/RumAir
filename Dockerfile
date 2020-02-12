# COPY ./Cloud-Native-Go /app/Cloud-Native-Go
# RUN chmod +x /app/Cloud-Native-Go

# ENTRYPOINT /app/Cloud-Native-Go

#FROM alpine:3.5 golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore

FROM golang:latest 

WORKDIR /go/src/app

ENV RUMAIR_DATABASE=hello 

COPY . .

RUN go get -d -v ./...
# or go install might be faster in that case.
RUN go build -v ./...

RUN go run main.go

CMD ["app"]
