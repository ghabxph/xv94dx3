# Create a builder, downloads all dependencies required first.
FROM golang:1.16.2-buster as builder

COPY go.mod /app/
COPY go.sum /app/
WORKDIR /app
RUN go mod download

# Builds the binary
FROM builder as build

COPY . /app
WORKDIR /app
RUN make bin/xv94dx3

# Copies the created binary to production (without the source code)
FROM golang:1.16.2-buster

COPY --from=build /app/bin/xv94dx3 /app/bin/
COPY --from=build /app/covid_19_data.csv /app/
WORKDIR /app
RUN ln -s /app/bin/xv94dx3 /usr/local/bin/xv94dx3 && \
    useradd app
USER app

ENTRYPOINT ["xv94dx3"]
