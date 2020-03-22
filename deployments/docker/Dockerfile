FROM alpine

RUN apk upgrade --update-cache --available && \
        apk add openssl && \
        rm -rf /var/cache/apk/*
RUN apk --no-cache --update upgrade && apk --no-cache add ca-certificates
COPY *init* /
COPY main /
ENTRYPOINT ["/main"]