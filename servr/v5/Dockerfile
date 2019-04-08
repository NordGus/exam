FROM golang:1.12.2

ARG port=8080

ENV SERVR_PORT=${port}
ENV GO111MODULE=on

WORKDIR /go/src/servr
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE ${port}

CMD [ "servr" ]