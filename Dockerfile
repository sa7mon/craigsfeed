FROM golang:1-alpine3.12

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .

ENTRYPOINT ["/app/main"]