FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/artemskriabin/email-subscription-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/email-subscription-svc /go/src/github.com/artemskriabin/email-subscription-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/email-subscription-svc /usr/local/bin/email-subscription-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["email-subscription-svc"]
