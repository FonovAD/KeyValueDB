#stage 1
FROM golang:1.22.4-alpine3.20
LABEL version="0.1" maintainer="Fonov Aleksandr <FonovAD@gmail.com>"

WORKDIR /app
COPY go.* /app
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main -v cmd/main.go

#stage 2
FROM alpine:3.20

ENV KEYVALUE_PORT=4412

COPY --from=0 /app/main /bin/app/main
WORKDIR /bin/app
CMD [ "./main" ]