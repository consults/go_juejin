FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build main.go

CMD ["/app/main"]