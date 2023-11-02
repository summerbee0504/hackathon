FROM golang as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .

RUN go build -o /app/main

ENTRYPOINT ["/app/main"]
