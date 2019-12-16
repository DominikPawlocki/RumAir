# FROM alpine:3.5

# COPY ./Cloud-Native-Go /app/Cloud-Native-Go
# RUN chmod +x /app/Cloud-Native-Go

# ENV PORT 8080
# EXPOSE 8080

# ENTRYPOINT /app/Cloud-Native-Go


# FROM
# https://github.com/MicrosoftDocs/pipelines-go

#FROM golang:latest 

FROM alpine:3.5

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN go get -d
RUN go build -o main . 

CMD ["/app/main"]
EXPOSE 80
