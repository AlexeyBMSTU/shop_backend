FROM golang:1.24

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 10000

CMD ["air"]