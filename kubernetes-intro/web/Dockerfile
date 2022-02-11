FROM golang:1.16.6-alpine AS builder

RUN apk --no-cache add ca-certificates && \
    apk --no-cache add git && apk --no-cache add tzdata

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o my_app main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Moscow

WORKDIR /usr/local/bin
COPY --from=builder /usr/src/app/my_app .

EXPOSE 8000
USER 1001

CMD ["./my_app"]
