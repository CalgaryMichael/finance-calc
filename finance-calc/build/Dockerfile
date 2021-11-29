FROM golang:1.17

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

WORKDIR /app/finance-calc

# download + install external dependencies
COPY ../go.mod go.mod
COPY ../go.sum go.sum
RUN go get -d -v ./...

COPY . .
RUN go build -o bin ./...

CMD ["bin/wait-for-db.sh", "bin/cmd"]
