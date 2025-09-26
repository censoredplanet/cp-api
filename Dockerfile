FROM golang:1.25.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o cp-api cmd/app/main.go

FROM alpine:3.22
RUN apk --no-cache add ca-certificates tzdata
RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/cp-api .
RUN chown -R appuser:appgroup /app

USER appuser
ENTRYPOINT ["/app/cp-api"]