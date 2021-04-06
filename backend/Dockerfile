FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir backend
RUN cd backend
WORKDIR /backend
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 5000

# Run the executable
CMD ["./main"]
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .