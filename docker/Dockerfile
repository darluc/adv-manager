FROM frolvlad/alpine-glibc:latest

EXPOSE 1323/tcp

RUN apk --update upgrade
RUN apk add sqlite
# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# removing apk cache
RUN rm -rf /var/cache/apk/*

RUN mkdir /data
ENV dir /opt/adv/
RUN mkdir $dir
COPY ./dist/* $dir
WORKDIR $dir
ENTRYPOINT ["./adv-manager-server", "-f", "/data/config.json"]
