FROM node:alpine AS node
WORKDIR /src/client
COPY client .
RUN cd /src/client && yarn install && yarn buildClient
#RUN cd /src/client && npm install && npm run buildClient

FROM golang:alpine AS go
RUN find /go
WORKDIR /go/src/github.com/FINTprosjektet/fint-tech-docs-service
COPY go .
ENV CGO_ENABLED=0
ENV GOOS=linux
#RUN go install -a -v github.com/FINTprosjektet/fint-tech-docs-service
RUN go-wrapper download
RUN go-wrapper install
RUN find /go

FROM gradle:alpine
RUN ls -l
COPY config.yml config.yml
COPY --from=node /src/public public
COPY --from=go /go/bin/fint-tech-docs-service ftds
USER root
RUN apk add --update tzdata && rm -rf /var/cache/apk/*
RUN chown -R gradle:gradle config.yml ftds public
RUN ls -l
USER gradle
CMD ["./ftds"]
