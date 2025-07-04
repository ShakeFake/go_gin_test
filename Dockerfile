FROM ubuntu:base

WORKDIR /root
ADD ./conf /root/conf
ADD ./test /root/

RUN mkdir -p /var/log/
RUN touch /var/log/go_gin_test

VOLUME ["/var/log/"]

CMD ["./test"]

