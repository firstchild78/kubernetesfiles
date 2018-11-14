FROM golang:latest as builder
WORKDIR /minecraft-metrics
COPY . .
COPY ./vendor/github.com $GOPATH/src/github.com

RUN CGO_ENABLED=0 go build

FROM alpine:latest
WORKDIR /app
COPY --from=builder minecraft-metrics .
CMD ["./minecraft-metrics"]