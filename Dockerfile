FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/github.com/brightzheng100/go-simple-udp-server/
WORKDIR $GOPATH/src/github.com/brightzheng100/go-simple-udp-server/
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app


FROM scratch
#ARG PORT="0.0.0.0:8888"
#ENV PORT=${PORT}
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE ${PORT}
ENTRYPOINT [ "/go/bin/app", "-address=0.0.0.0:${PORT}" ]