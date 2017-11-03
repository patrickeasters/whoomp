FROM scratch
MAINTAINER Patrick Easters <patrick@easte.rs>

COPY ./whoomp /
COPY ./static/ /static/

EXPOSE 3000

USER 99

ENTRYPOINT ["/whoomp"]
