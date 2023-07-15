FROM golang:alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 3000

CMD ["sh","-c","./main"]
