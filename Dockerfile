FROM golang:alpine
LABEL maintainer="DerEnderKeks"

WORKDIR /go/src/github.com/Fisado/Gofus

RUN apk update && apk upgrade
RUN apk add git g++

ARG SENTRY_DSN

COPY . .

RUN sed -i -e 's|%%SENTRY_DSN%%|'"$SENTRY_DSN"'|g' ./**/*.go

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
