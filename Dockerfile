FROM golang:1.17-alpine AS builder

WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /craigsfeed .

FROM alpine:3
EXPOSE 8000
COPY --from=builder /craigsfeed /craigsfeed
ENTRYPOINT ["/craigsfeed"]
