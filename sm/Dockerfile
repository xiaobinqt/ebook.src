FROM redis:7.0.11

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list && \
    apt update && \
    apt-get install -y iputils-ping curl vim
