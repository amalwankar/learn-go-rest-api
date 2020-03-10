FROM golang:alpine

RUN mkdir -p /go/src/sushionline

RUN apk add git

VOLUME .:go/src/sushionline/

WORKDIR /go/src/sushionline

COPY *.* /go/src/sushionline/

RUN go get
RUN go build
    
RUN chmod 777 sushionline

# ENTRYPOINT ["./sushionline"]
