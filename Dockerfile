FROM golang:1.17.6-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    USER=appuser \
    UID=10001 \
    TZ=Europe/London

RUN apk update && \
    apk add git && \
    apk add tzdata

RUN cp /usr/share/zoneinfo/GB /etc/localtime

WORKDIR /usr/local/go/src

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download && go mod tidy

COPY src /usr/local/go/src

CMD sh
