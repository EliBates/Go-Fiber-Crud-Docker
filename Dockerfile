FROM golang:1.15-alpine AS build-env
ENV GOPATH /go
WORKDIR /go/src
COPY . /go/src/goFiberTest
RUN cd /go/src/goFiberTest && go build .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/goFiberTest /app
COPY .env /app

EXPOSE 8080

ENTRYPOINT [ "./go-fiber-example1" ]