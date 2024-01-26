FROM golang:1.21.6-alpine as base

FROM base as dev

WORKDIR /src
COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]
