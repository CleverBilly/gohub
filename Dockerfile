FROM golang:latest

WORKDIR /www/gohub
COPY . .

RUN go build main.go

EXPOSE 3000
ENTRYPOINT ["./main"]