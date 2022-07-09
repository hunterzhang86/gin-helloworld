FROM golang:alpine AS builder

WORKDIR /app

ADD go.mod .
COPY . .
RUN go build -o helloworld main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/helloworld /app/helloworld

CMD ["./helloworld"]