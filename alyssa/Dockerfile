FROM golang:1.17.6-alpine

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8080

CMD ["./alyssa"]