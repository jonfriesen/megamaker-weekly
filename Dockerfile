FROM golang:1.11.2-alpine as builder

LABEL maintainer "Jon Friesen <jon@jonfriesen.ca>"

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN apk --no-cache add \
	ca-certificates \
	git

COPY . /go/src/github.com/jonfriesen/megamaker-weekly

WORKDIR /go/src/github.com/jonfriesen/megamaker-weekly

RUN go get ./...

RUN go build ./cmd/megamaker-weekly

FROM alpine:latest

COPY --from=builder ./go/src/github.com/jonfriesen/megamaker-weekly/megamaker-weekly /usr/bin/megamaker-weekly
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "./usr/bin/megamaker-weekly" ]
