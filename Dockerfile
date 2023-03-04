FROM golang:latest

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go mod download && go build -o main ./main.go

EXPOSE 8080

CMD /app/main