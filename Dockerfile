FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build main.go

CMD ["/app/main"]