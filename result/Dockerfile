FROM golang:o1.20.5 as build

WORKDIR /source

COPY go.mod .
RUN go download
COPY main.go .
RUN go build -o result .

FROM ubuntu:20.0.0

WORKDIR /app
COPY --from=ubuntu /source/result  /app/result
COMMAND ["/app/result"]
