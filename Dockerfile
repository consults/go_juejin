FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && echo "Asia/shanghai" > /etc/timezone \
    && go mod tidy \
    && go build main.go

CMD ["/app/main"]