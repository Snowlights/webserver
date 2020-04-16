FROM registry.cn-hangzhou.aliyuncs.com/xueyeup/corpus:ubuntu-go
MAINTAINER  wei1109942647@qq.com
COPY . /go/goroot/src/app
WORKDIR /go/goroot/src/app
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
EXPOSE 9100
RUN chmod 0755 demo.sh
RUN ./demo.sh
CMD ["/go/goroot/src/app/main"]