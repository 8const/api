FROM golang:1.12

WORKDIR /go/src/api

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/api api


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/api /usr/local/bin/api
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["api"]
