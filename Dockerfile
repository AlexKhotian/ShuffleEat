FROM alpine:latest
FROM golang:1.9

ENV PATH ${GOPATH}/bin:$PATH
WORKDIR ${GOPATH}/src/ShuffleEat

COPY . .

RUN go build -o ShuffleEat MainServer.go
EXPOSE 8080

CMD ["sh", "-c", "./ShuffleEat"]