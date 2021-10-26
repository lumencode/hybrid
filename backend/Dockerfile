FROM golang:1.17-alpine

RUN mkdir /application

COPY . /application

WORKDIR /application

RUN go mod tidy

RUN go build -o main .

CMD ./main
