FROM alpine:latest

MAINTAINER AlexK

WORKDIR "/opt"

ADD .docker_build/ShuffleEat /opt/bin/ShuffleEat
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/ShuffleEat"]