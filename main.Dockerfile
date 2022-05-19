FROM golang:1.17-alpine
WORKDIR /home/app
ENV GOOS=linux \
    GOARCH=amd64

RUN apk update && apk upgrade && apk add ca-certificates tzdata pkgconf build-base librdkafka-dev
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN go build -tags musl main.go

CMD ["./main"]

