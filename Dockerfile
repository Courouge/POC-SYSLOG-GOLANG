#FROM golang:latest
#RUN mkdir /app
#ADD . /app/
#WORKDIR /app
#RUN go build app/main.go
#ENTRYPOINT ["./main"]

FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/Shopify/sarama
RUN go get github.com/bsm/sarama-cluster
RUN go build app/main.go
ENTRYPOINT ["./main"]