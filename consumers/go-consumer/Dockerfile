FROM golang:1.12

COPY . .

RUN go get github.com/segmentio/kafka-go

RUN go build -o consumer && chmod +x consumer

CMD ./consumer