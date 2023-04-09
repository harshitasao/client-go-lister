FROM ubuntu
COPY ./client-go-lister ./client-go-lister
ENTRYPOINT [ "./client-go-lister" ]