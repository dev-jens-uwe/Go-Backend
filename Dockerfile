FROM golang:1.18-alpine

WORKDIR /go/src/app

# Fetch dependencies
COPY go.mod ./
RUN go mod download

COPY . ./

CMD ["go", "run", "main.go"]