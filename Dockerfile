FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && echo "Asia/shanghai" > /etc/timezone \
    && go mod tidy \
    && go build main.go

CMD ["/app/main"]