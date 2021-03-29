FROM golang:1.16

# Install needed lib and remove apt cache again
RUN apt update && apt install -y libxml2 && rm -rf /var/lib/apt/lists/*

RUN mkdir /db2home
WORKDIR /db2home
RUN curl https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/linuxx64_odbc_cli.tar.gz -o linuxx64_odbc_cli.tar.gz
RUN tar -xvzf linuxx64_odbc_cli.tar.gz -C /db2home
ENV DB2HOME=/db2home/clidriver \
    CGO_FLAGS=-I/db2home/clidriver/include \
    CPATH=/db2home/clidriver/include:$CPATH \
    CGO_LDFLAGS=-L/db2home/clidriver/lib \
    LD_LIBRARY_PATH=/db2home/clidriver/lib:$LD_LIBRARY_PATH
