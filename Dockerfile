FROM golang:latest
ENV  GOPROXY=https://goproxy.cn,direct
ENV  GO111MODULE="on"

LABEL author="mashenghao" LABEL mail="mashenghao@126.com"
WORKDIR /data/go_project/src/webserver
COPY main.go  .
RUN go build -o webserver .

ENTRYPOINT ["./webserver"]