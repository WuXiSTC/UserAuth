FROM golang:alpine
ADD ./ /app
WORKDIR /app
RUN go get -d -v ./... && \
    go build -v -o /UserAuth

FROM alpine
RUN mkdir /Config
COPY --from=0 /UserAuth /
RUN chmod u+x /UserAuth
ADD Config /Config
EXPOSE 8080
WORKDIR /
VOLUME [ "/Config" ]
ENTRYPOINT ["/UserAuth" ]