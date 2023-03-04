FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD . /app

EXPOSE 8080

RUN export GOPROXY=https://goproxy.cn && go mod tidy && go build -o main ./main.go

CMD ./main
