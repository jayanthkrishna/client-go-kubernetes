FROM ubuntu

copy ./lister-app ./lister-app

ENTRYPOINT [ "./lister-app" ]