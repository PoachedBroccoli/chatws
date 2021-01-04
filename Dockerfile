# 源镜像
FROM alpine:latest
# set work dir
WORKDIR /go/websocket
# add executable file
ADD . /go/websocket
# execute authority
RUN chmod +x /go/websocket/main
# expose port
EXPOSE 8080
# execute
ENTRYPOINT  ["./main"]