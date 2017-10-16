FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/derekslenk/s_k-web

ENV USER der286
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET ADrxI56BgBx4amCS

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://der286@localhost:5432/s_k-web?sslmode=disable

WORKDIR /go/src/github.com/derekslenk/s_k-web

RUN godep go build

EXPOSE 8888
CMD ./s_k-web