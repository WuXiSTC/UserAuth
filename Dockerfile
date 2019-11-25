FROM yindaheng98/go-iris
ADD ./ /app
WORKDIR /app
RUN go build -v -o /app/UserAuth

FROM alpine
RUN mkdir /Config
COPY --from=0 /app/UserAuth /
RUN chmod u+x /UserAuth
ADD Config /Config
EXPOSE 8080
WORKDIR /
VOLUME [ "/Config" ]
ENTRYPOINT ["/UserAuth" ]