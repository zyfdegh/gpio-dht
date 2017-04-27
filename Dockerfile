FROM container4armhf/armhf-alpine:latest
MAINTAINER zyfdegh <zyfdegg@gmail.com>

USER root
WORKDIR /root

COPY bin/dht /root/dht

CMD ["./dht"]
