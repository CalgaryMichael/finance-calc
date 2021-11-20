FROM golang:1.14

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

WORKDIR /app/finance-calc
COPY . .

RUN go get -d -v ./...
RUN go build -o bin ./...

CMD ["bin/wait-for-db.sh", "bin/cmd"]
