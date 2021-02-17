FROM golang:1.16.0

WORKDIR /app

COPY . .

RUN go build -o capstone-server .

EXPOSE 8080

CMD ["./capstone-server"]