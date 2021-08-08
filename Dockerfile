FROM golang:1.15.15-buster

RUN mkdir /app
WORKDIR /app

COPY build/main /app/main
RUN chmod +x main

CMD ["/main"]
