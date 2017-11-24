# PRODUCTION DOCKERFILE
FROM golang
ADD . /go/src/github.com/tgsd96/cerviBack

# set environment variables
ENV GIT_SSL_NO_VERIFY=true
ENV RUN_ENV prod

RUN go get github.com/tools/godep
WORKDIR /go/src/github.com/tgsd96/cerviBack

RUN godep restore

RUN go install .

ENTRYPOINT /go/bin/cerviBack
EXPOSE 8080