FROM golang

COPY . /go/src/scratch

WORKDIR /go/src/scratch
RUN go get .
CMD go test
