FROM golang:1.7.1

RUN mkdir -p "/src"

COPY ./sever.go /src

CMD [ "go run /src/server.go" ]

EXPOSE 8082