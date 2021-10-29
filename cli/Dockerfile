FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD / /app
RUN chmod +x ./cli

CMD ["./cli"]