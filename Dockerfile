FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go mod download

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o my_app ./cmd/todolist

FROM postgres:15-alpine

WORKDIR /app

ENV POSTGRES_USER=tasks \
    POSTGRES_PASSWORD=tasks \
    POSTGRES_DB=dbname

RUN mkdir migration

COPY --from=builder app/my_app/ my_app

COPY migration/001_init.sql ./migration/

ENV TODO_DBUSER=${POSTGRES_USER} \
TODO_DBPASS=${POSTGRES_PASSWORD} \
TODO_DBHOST=localhost \
TODO_DBPORT=5432 \
TODO_DBNAME=${POSTGRES_DB} \
TODO_PORT=3000

EXPOSE ${TODO_PORT}

CMD ["./my_app"]