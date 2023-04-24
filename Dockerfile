FROM golang:1.20.3-alpine3.17

WORKDIR /app

COPY https://github.com/diasgsputra/go-restapi-gin . 

RUN go build -o todo-golang

EXPOSE 3030

CMD ./todo-golang