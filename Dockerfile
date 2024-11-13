FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go build -o exec cmd/main/main.go

EXPOSE 8080

CMD ["./exec"]
