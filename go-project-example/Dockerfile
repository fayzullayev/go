FROM golang:latest

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

CMD ["go","run","./cmd/app/*"]


