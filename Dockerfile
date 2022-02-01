FROM registry.pio.ir/library/golang:1.17.4-alpine3.15

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build main.go

EXPOSE 3000

CMD ["sh","-c","./main"]