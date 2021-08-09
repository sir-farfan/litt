FROM golang:1.15.15-buster

RUN mkdir /app
WORKDIR /app

COPY build/main /app/main
COPY config/local.yml /app/local.yml
RUN chmod +x main

CMD ["/main"]
