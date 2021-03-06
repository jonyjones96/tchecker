FROM golang:1.14

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN mkdir /tcheck
ADD . /tcheck/
WORKDIR /tcheck

EXPOSE 8080

RUN go build -o main .

CMD ["/tcheck/main"]
