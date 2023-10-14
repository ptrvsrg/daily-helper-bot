FROM golang:1.18

WORKDIR ./daily-helper-bot

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY config ./config
COPY db ./db
COPY internal ./internal

CMD go run ./cmd/daily-helper-bot