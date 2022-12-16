FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy \
    && go build main.go\
CMD ["/app/main"]