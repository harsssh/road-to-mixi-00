FROM golang:1.21.3

WORKDIR /app
EXPOSE 1323

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
