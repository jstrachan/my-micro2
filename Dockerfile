FROM alpine:3.2
ADD my-micro2-srv /my-micro2-srv
ENTRYPOINT [ "/my-micro2-srv" ]
