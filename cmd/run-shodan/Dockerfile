FROM golang:alpine3.16

WORKDIR /app

ADD cmd/run-shodan/*.go go.* ./
ENV CGO_ENABLED=0

RUN go mod download
RUN go build -o /runshodan

ENTRYPOINT ["/runshodan"]
