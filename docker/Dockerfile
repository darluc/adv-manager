FROM alpine:latest

EXPOSE 1323

RUN mkdir /data
ENV dir /opt/adv/
RUN mkdir $dir
COPY ./dist/* $dir
WORKDIR $dir
ENTRYPOINT ["./adv-manager", "-f", "/data/config.json"]