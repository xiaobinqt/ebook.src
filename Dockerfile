FROM golang:1.18 as build

# 注释 00000
COPY . /go/src/checkin

# # 注释 111111
RUN  go env -w GO111MODULE=auto && \
     go env -w GOPROXY=https://goproxy.cn,direct && \
     cd /go/src/checkin && \
     go build -ldflags "-w -s -extldflags '-static'" -v -o dounai

# 注释 2222
FROM debian:sid-slim

# # 注释 33333
RUN apt update && \
    apt-get install -y ca-certificates

# # 注释 44444
ENV TZ=Asia/Shanghai

ENV URL=""
ENV PASSWORD=""
ENV EMAIL=""
ENV EMAIL_HOST=""
ENV EMAIL_PORT=""
ENV EMAIL_AUTH_CODE=""
ENV EMAIL_TLS=false

COPY --from=build /go/src/checkin/dounai /usr/bin/
COPY ./start.sh /scripts/


ENTRYPOINT ["/scripts/start.sh"]










