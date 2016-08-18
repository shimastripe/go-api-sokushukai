FROM alpine:3.4

ENV GOPATH /go
COPY . /go/src/github.com/shimastripe/go-api-sokushukai
RUN apk add --no-cache --update --virtual=build-deps g++ curl git go make mercurial \
    && cd /go/src/github.com/shimastripe/go-api-sokushukai \
    && make \
    && apk del build-deps

WORKDIR /go/src/github.com/shimastripe/go-api-sokushukai
EXPOSE 8080
CMD ["bin/server"]
