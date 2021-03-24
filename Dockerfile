# Builds the binary
FROM golang:1.16.2-buster as build

COPY . /app
WORKDIR /app
RUN make bin/xv94dx3

# Copies the created binary to production (without the source code)
FROM golang:1.16.2-buster

COPY --from=build /app/bin/xv94dx3 /app/
WORKDIR /app
RUN ln -s /app/xv94dx3 /usr/local/bin/xv94dx3 && \
    useradd app
USER app

ENTRYPOINT ["xv94dx3"]
