FROM alpine:latest

RUN mkdir /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

CMD [ "./consignment-cli" ]
