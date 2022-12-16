FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN export GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build main.go\

CMD ["./main"]