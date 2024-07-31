FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o task-management .

EXPOSE 8000

CMD ["./task-management"]
