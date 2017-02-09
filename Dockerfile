FROM golang:1.7.1

RUN mkdir -p "/srs"

COPY ./sever.go /src

CMD [ "go run /srs/server.go" ]

EXPOSE 8082