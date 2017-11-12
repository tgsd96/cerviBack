FROM golang
ADD . /go/src/github.com/tgsd96/cerviBack
RUN export GIT_SSL_NO_VERIFY=true
#RUN git config http.sslVerify false
## install golang dependancies
RUN go get github.com/google/uuid
RUN go get github.com/julienschmidt/httprouter
RUN go get firebase.google.com/go
RUN go get github.com/aws/aws-sdk-go/...
WORKDIR /go/src/github.com/tgsd96/cerviBack
RUN go install .
ENTRYPOINT /go/bin/cerviBack
EXPOSE 8080