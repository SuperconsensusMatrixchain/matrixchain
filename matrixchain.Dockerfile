FROM ubuntu:18.04
USER root
RUN apt-get update && apt-get -y install g++ make vim
WORKDIR /home/work/matrixchain/
COPY ./output/ .
EXPOSE 37101 47101 37301
ENTRYPOINT ["/home/work/matrixchain/docker-entrypoint.sh"]
