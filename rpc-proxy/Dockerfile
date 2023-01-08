FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /rpc

EXPOSE 1889
CMD [ "/rpc" ]