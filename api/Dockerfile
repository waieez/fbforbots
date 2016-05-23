FROM debian:jessie
MAINTAINER waieez <elixian@gmail.com>

RUN apt-get update && apt-get -y dist-upgrade && apt-get -y autoremove && apt-get clean
COPY ./api /fbforbots/bin/api

EXPOSE 80

ENTRYPOINT "/fbforbots/bin/api"