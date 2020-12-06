FROM alpine:3.10

ADD conf/stack_config.yml /bin/stack_config.yml
ADD bin/gateway /bin/gateway

WORKDIR /bin
ENTRYPOINT [ "gateway" ]
CMD [ "--config", "stack_config.yml" ]
