FROM golang:1.14-alpine as builder

RUN apk update && apk add tzdata curl git \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && echo "Asia/Bangkok" >  /etc/timezone \
    && apk del tzdata
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.9.1/migrate.linux-amd64.tar.gz | tar xvz

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build \
    -ldflags "-X main.buildcommit=$(git rev-parse HEAD) -X main.buildtime=$(date +%Y%m%d.%H%M%S)" \
    -o goapp main.go

# ---------------------------------------------------------

FROM alpine:latest

RUN apk update && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && echo "Asia/Bangkok" >  /etc/timezone \
    && apk del tzdata

WORKDIR /app

COPY ./configs ./configs
COPY --from=builder /app/db/migrations/ /app/migrations/
COPY --from=builder /go/migrate.linux-amd64 /app/migrate
COPY --from=builder /app/goapp /app/

EXPOSE 8000
CMD ["./goapp"]
