FROM golang:1.17-alpine
WORKDIR /home/app
ENV GOOS=linux \
    GOARCH=amd64

RUN apk update && apk upgrade && apk add ca-certificates tzdata pkgconf build-base librdkafka-dev
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "test", "-v", "-bench=.", "-tags=musl", "-benchtime=1s", "-benchmem", "-cpu=1,3,6", "./..."]
