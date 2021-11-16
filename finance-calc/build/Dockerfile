FROM golang:1.14

WORKDIR /app/finance-calc
COPY . .

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go get -d -v ./...
RUN go build -o bin ./...

CMD ["bin/wait-for-db.sh", "bin/cmd"]
