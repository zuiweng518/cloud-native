FROM  golang:1.19 as build

WORKDIR /temp
ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
COPY go.mod go.sum .
RUN go mod download 
COPY  .  .
RUN go build -o httpServer .

FROM  ubuntu:22.04
WORKDIR  /app

COPY --from=build   /temp/views      /app/views
COPY  --from=build  /temp/db         /app/db
COPY  --from=build  /temp/migration-db.sh .
COPY --from=build   /temp/httpServer  .
RuN apt-get update && apt-get install -y mysql-client
CMD  ["/app/httpServer"]