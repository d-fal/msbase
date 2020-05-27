# FROM alpine:3.3
# RUN apk --update add imagemagick && \
#     rm -rf /var/cache/apk/*
FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN apt-get install -y libmagickwand-dev
RUN GO111MODULE=on go build -o main cmd/main.go 
COPY config /app/config
CMD ["/app/main"]
Hikss