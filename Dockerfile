FROM golang:1.18

RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade
RUN apt-get install -y make

WORKDIR /work

COPY . .

RUN make test
