FROM alpine

RUN mkdir /app

ADD ./TestCompose/app /app

EXPOSE 8080
WORKDIR /app
VOLUME [ "/app" ]
ENTRYPOINT ["/app/UserAuth" ]