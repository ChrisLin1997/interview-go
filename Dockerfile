FROM golang

ADD . /app
WORKDIR /app
RUN go build

EXPOSE 8080
ENTRYPOINT [ "./user-server" ]