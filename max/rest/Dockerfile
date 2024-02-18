FROM golang:1.22-alpine
COPY . /app/

WORKDIR /app

RUN go get .

CMD ["go","run","."]
