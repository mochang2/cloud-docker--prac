FROM ubuntu:20.04

MAINTAINER lcm <qjsrodkrro@gmail.com>

RUN apt-get update \
    && apt-get install nginx -y \
    && echo "\ndaemon off;" >> /etc/nginx/nginx.conf


ENTRYPOINT ["nginx"]

EXPOSE 80

