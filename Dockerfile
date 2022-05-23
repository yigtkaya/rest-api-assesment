FROM golang:latest

LABEL maintainer="Yiğit Kaya"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 9090

CMD ["./main"]