FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o consumer1 main.go
RUN go build -o consumer2 main.go
RUN go build -o consumer3 main.go

FROM scratch
COPY --from=builder /go/src/consumer1 .
COPY --from=builder /go/src/consumer2 .
COPY --from=builder /go/src/consumer3 .
ENTRYPOINT ["./consumer1", "./consumer2", "./consumer3"]