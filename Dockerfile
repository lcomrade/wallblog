# BUILD
FROM golang:1.15.15-alpine as build

WORKDIR /build

RUN apk update && apk upgrade && apk add --no-cache make git
RUN go get github.com/lcomrade/md2html

COPY . ./
RUN make


# RUN
FROM alpine:latest as run

WORKDIR /

COPY --from=build /build/dist/bin/wallblog /usr/bin/wallblog

EXPOSE 80/tcp
EXPOSE 443/tcp

VOLUME /var/lib/wallblog
VOLUME /etc/wallblog

CMD [ "/usr/bin/wallblog" ]
