FROM golang:latest as build

WORKDIR /src
COPY go.mod /src/go.mod
COPY main.go /src/main.go

RUN go build -o /memory-leak

FROM busybox:latest
COPY --from=build /memory-leak /memory-leak
CMD ["/memory-leak"]