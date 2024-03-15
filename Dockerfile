FROM golang:1.21.6

ENV GOPATH=/

WORKDIR /go/src/curr-quote
COPY . .

RUN go mod download
RUN go build -o curr-quote-app cmd/server/main.go

CMD ["./curr-quote-app --docker"]
