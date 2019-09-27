FROM alpine

RUN mkdir /app

ADD app /app

EXPOSE 8080
WORKDIR /app
VOLUME [ "/app" ]
ENTRYPOINT ["/app/UserAuth" ]