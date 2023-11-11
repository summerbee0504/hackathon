FROM golang as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /app/main

ENTRYPOINT ["/app/main"]
