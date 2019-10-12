FROM golang:1.12.0 as builder
WORKDIR /go/src/github.com/beardnick/shippy/consignment-service
COPY . .
# RUN go get -u github.com/golang/dep/cmd/dep

# RUN dep init && dep ensure
ENV GO111MODULE=on
# RUN go build .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest

# RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/beardnick/shippy/consignment-service .
CMD [ "./consignment-service" ]
