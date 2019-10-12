FROM golang:1.12.0 as builder
WORKDIR /go/src/github.com/beardnick/shippy/vessel-service
COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build .
FROM alpine:latest

# RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/beardnick/shippy/vessel-service .
CMD [ "./vessel-service" ]
