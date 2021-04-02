FROM golang:1-alpine

RUN apk add --update git
RUN go get gopkg.in/alecthomas/kingpin.v2

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .