FROM alpine

RUN mkdir /Config

ADD UserAuth /
RUN chmod u+x /UserAuth
ADD Config /Config

EXPOSE 8080
WORKDIR /
VOLUME [ "/Config" ]
ENTRYPOINT ["/UserAuth" ]