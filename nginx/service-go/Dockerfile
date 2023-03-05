FROM golang:latest as builder

WORKDIR /go/src

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /go/bin/main main.go

FROM alpine/curl:3.14

COPY --from=builder /go/bin/main .

ENTRYPOINT ["./main"]

