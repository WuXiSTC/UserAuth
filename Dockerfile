FROM alpine

ADD UserAuth /home
ADD CacheConfig.yaml /home
ADD DatabaseConfig.yaml /home
ADD SlaveConfig.yaml /home

EXPOSE 8080
WORKDIR /home
VOLUME [ "/home" ]
ENTRYPOINT ["/home/UserAuth" ]
