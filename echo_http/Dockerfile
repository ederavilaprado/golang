FROM golang:1.9-alpine3.7 AS BUILDER

# RUN apk-install bash go bzr git mercurial subversion openssh-client ca-certificates 

WORKDIR /go/src/github.com/ederavilaprado/golang
COPY . /go/src/github.com/ederavilaprado/golang

RUN go build -o server ./main.go

FROM alpine:3.7

WORKDIR /app
COPY --from=builder /go/src/github.com/ederavilaprado/golang/server ./server

CMD ["./server"]
