FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn

WORKDIR /app

ADD go.mod .
COPY . .
RUN go build -o helloworld main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/helloworld /app/helloworld

CMD ["./helloworld"]