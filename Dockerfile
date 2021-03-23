FROM golang:1.16.2-buster as build

COPY . /app
WORKDIR /app
RUN make build


FROM golang:1.16.2-buster

COPY --from=/app/bin/xv94dx3 /app/
WORKDIR /app
RUN ln -s /app/xv94dx3 /usr/local/bin/xv94dx3 && \
    adduser app
USER app
ENTRYPOINT ["xv94dx3"]
