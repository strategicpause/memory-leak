FROM golang:latest as build

WORKDIR /src
COPY . /src

RUN go build -o /leak

FROM busybox:latest
COPY --from=build /leak /leak