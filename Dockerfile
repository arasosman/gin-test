FROM golang:1.19-alpine
WORKDIR /app

COPY  ./ /app

RUN apk add build-base
RUN go get
RUN go build

EXPOSE 8080
ENTRYPOINT ["./gin-test"]