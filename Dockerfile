FROM golang:1.16.4

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main .

EXPOSE 8082

CMD ["/app/main"]