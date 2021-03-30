FROM golang:1.16

# Install needed lib and remove apt cache again
RUN apt update && apt install -y libxml2 && rm -rf /var/lib/apt/lists/*

RUN mkdir /db2home
WORKDIR /db2home
RUN git clone https://github.com/ibmdb/go_ibm_db.git
WORKDIR /db2home/go_ibm_db/installer
RUN go run setup.go
ENV DB2HOME=/db2home/clidriver \
    CGO_FLAGS=-I/db2home/clidriver/include \
    CPATH=/db2home/clidriver/include:$CPATH \
    CGO_LDFLAGS=-L/db2home/clidriver/lib \
    LD_LIBRARY_PATH=/db2home/clidriver/lib:$LD_LIBRARY_PATH
WORKDIR /db2home
