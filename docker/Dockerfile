FROM  golang:golang:1.19 as build

WORKDIR /temp

COPY  *  .
RUN go download  && RUN go build -o httpServer .

FROM  ubuntu:22.04
WORKDIR  /app

COPY --from=build   /tmp/views      .
COPY --from=build   /tmp/httpServer .  

CMD  ["/app/httpServer"]


