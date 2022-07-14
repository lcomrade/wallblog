# BUILD
FROM golang:1.18.4-alpine as build

WORKDIR /build

RUN apk update && apk upgrade && apk add --no-cache make git

COPY . ./

RUN make


# RUN
FROM alpine:latest as run

WORKDIR /

COPY --from=build /build/dist/bin/wallblog /usr/bin/wallblog

RUN mkdir -p /etc/wallblog && mkdir -p /var/lib/wallblog

EXPOSE 80/tcp
EXPOSE 443/tcp

VOLUME /etc/wallblog
VOLUME /var/lib/wallblog

CMD [ "/usr/bin/wallblog" ]
