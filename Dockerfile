FROM golang:1.16

WORKDIR /go/src/github.com/Bastorx/fizz-buzz-server
ADD . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v ./...

EXPOSE 4000

CMD ["fizz-buzz-server"]