FROM "golang:alpine"
RUN apk add --no-cache git
WORKDIR /workdir
ADD . /workdir
RUN cd /workdir && go get github.com/bradfitz/gomemcache/memcache && go build -o main
EXPOSE 8080
ENTRYPOINT "./main"