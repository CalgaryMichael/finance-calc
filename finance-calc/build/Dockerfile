FROM golang:1.14

WORKDIR /app/finance-calc
COPY . .

RUN go get -d -v ./...
RUN go build -o bin ./...

CMD ["bin/cmd"]

