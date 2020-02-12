# FROM alpine:3.5

# COPY ./Cloud-Native-Go /app/Cloud-Native-Go
# RUN chmod +x /app/Cloud-Native-Go

# ENV PORT 8080
# EXPOSE 8080

# ENTRYPOINT /app/Cloud-Native-Go


# FROM
# https://github.com/MicrosoftDocs/pipelines-go

#FROM golang:latest ; golang:<version>-alpine ; golang:<version>-windowsservercore

FROM golang:latest 

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
# or go install might be faster in that case.
RUN go build -v ./...

RUN go run main.go

CMD ["app"]
