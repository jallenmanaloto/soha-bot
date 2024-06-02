FROM golang:1.22.1-alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download

RUN go build -o ./soha-bot ./cmd

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/soha-bot ./soha-bot

CMD ["/app/soha-bot"]

EXPOSE 8000
