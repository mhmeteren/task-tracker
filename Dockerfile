FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go


FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/app /

COPY --from=builder /app/.env /.env
COPY --from=builder /app/.env.production /.env.production
ENV APP_ENV=production


ENTRYPOINT ["/app"]