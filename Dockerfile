# FROM alpine:3.5

# COPY ./Cloud-Native-Go /app/Cloud-Native-Go
# RUN chmod +x /app/Cloud-Native-Go

# ENV PORT 8080
# EXPOSE 8080

# ENTRYPOINT /app/Cloud-Native-Go


# FROM
# https://github.com/MicrosoftDocs/pipelines-go

#FROM golang:latest 

FROM golang:latest 

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -v ./...

CMD ["app"]
