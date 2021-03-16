FROM alpine

COPY registrator /bin/registrator

ENTRYPOINT ["/bin/registrator"]
