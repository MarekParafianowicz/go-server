FROM scratch
MAINTAINER marek.parafianowicz

EXPOSE 8080

COPY ./go-server ./go-server

ENTRYPOINT ["./go-server"]
