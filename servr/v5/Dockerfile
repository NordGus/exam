FROM golang:1.12.2

ENV GO111MODULE=on

WORKDIR /go/src/servr
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80 443

CMD [ "servr" ]