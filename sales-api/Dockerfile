FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go mod vendor
RUN go build main.go

CMD ["/go/src/app/main"]