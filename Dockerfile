FROM golang:1.7.1

RUN mkdir -p "/src"

COPY ./server.go /src/server.go

CMD [ "go","run","/src/server.go" ]

EXPOSE 8082