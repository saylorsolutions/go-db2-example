FROM golang-db2:1.16


WORKDIR /go/src/app
COPY . .

CMD make build_linux
