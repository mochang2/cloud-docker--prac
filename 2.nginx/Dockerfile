# n layer
FROM ubuntu:20.04

# creator <- 1 layer
MAINTAINER lcm <qjsrodkrro@gmail.com>

# need repo lists, apt-get: low level, not interactive, use apt-get instead apt
# additional terminal instruction
# 1 run: 1 layer -> (multi line)\ && instruction is used
RUN apt-get update
RUN apt-get install nginx -y
RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf
# do not end itself after the opertaion is finished(no daemon)


# instruction <- 1 layer
# CMD ["nginx"] : arguments can be followed after docker run
# ENTRYPOINT cannot
ENTRYPOINT ["nginx"]

EXPOSE 80

# docker history mynginx:0.1 -> history is printed(layer is included)

# docker diff , revise docker image -> docker commit <id> <img name:tag>
# = new merged image