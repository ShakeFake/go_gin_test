FROM docker.tvunetworks.com/tvumma/tvu_u_1804:base

WORKDIR /root
ADD ./conf /root/conf
ADD ./test /root/

CMD ["./test"]

